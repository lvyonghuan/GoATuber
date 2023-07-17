package gpt

import (
	"GoTuber/MEMORY"
	memory_gpt "GoTuber/MEMORY/NLPmodel/gpt"
	"GoTuber/NLP/service/gpt/function"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/config"
	"GoTuber/NLP/service/out"
	backend "GoTuber/frontend/live2d_backend"
	"GoTuber/proxy"
)

// GenerateTextByOpenAI 文本请求
func GenerateTextByOpenAI(msg *model.Msg) {
	var postDataTemp interface{}
	//函数调用流程
	if function.UseFunction {
		var postData postDataWithFunction
		postData.initRequestModel(msg)
		postDataTemp = postData
	} else { //为了健壮性
		var postData postDataWithoutFunction
		postData.initRequestModel(msg)
		postDataTemp = postData

	}
	//构造请求体
	postDataBytes, err := json.Marshal(postDataTemp)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
	}
	req, _ := http.NewRequest("POST", OpenAIChatUrl, bytes.NewBuffer(postDataBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GPTCfg.OpenAi.ApiKey)
	//发送请求
	client, err := proxy.Client()
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
		return
	}
	//处理返回
	defer resp.Body.Close()
	if resp == nil {
		backend.WebsocketToNLP <- true
		log.Println("response is nil")
		return
	}
	body, _ := io.ReadAll(resp.Body)
	var openAiRcv OpenAiRcv
	err = json.Unmarshal(body, &openAiRcv)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
		return
	}
	if len(openAiRcv.Choices) == 0 {
		backend.WebsocketToNLP <- true
		log.Println("OpenAI API调用失败，返回内容：", string(body))
		return
	}

	//检查是否使用函数调用
	if openAiRcv.Choices[0].FinishReason == "function_call" {
		openAiRcv = secondRequest(postDataTemp.(postDataWithFunction), openAiRcv)
		if reflect.ValueOf(openAiRcv).IsZero() {
			backend.WebsocketToNLP <- true
			return
		}
	}

	log.Println("openai生成内容：", openAiRcv.Choices[0].Message.Content)
	openAiRcv.Choices[0].Message.Content = strings.Replace(openAiRcv.Choices[0].Message.Content, "\n\n", "", 1)
	log.Printf("Model: %s TotalTokens: %d+%d=%d", openAiRcv.Model, openAiRcv.Usage.PromptTokens, openAiRcv.Usage.CompletionTokes, openAiRcv.Usage.TotalTokens)

	//压入AI的回答，形成短期记忆
	messagesAI := RequestMessages{
		Role:    "assistant",
		Content: openAiRcv.Choices[0].Message.Content,
		Name:    "AI",
	}
	MS = append(MS, messagesAI)

	if MEMORY.MemoryCfg.IsUse {
		memory := memory_gpt.Chat{
			Human: msg.Msg,
			AI:    openAiRcv.Choices[0].Message.Content,
		}
		memory.UserName = msg.Name
		memory.Type = "chat"
		memory.Namespace = "live"
		go memory.StoreMessage()
	}

	if openAiRcv.Usage.TotalTokens > 500 {
		//TODO:用户选择
		//cleanAllMessage()
		cleanFirstMessage()
	}
	var Msg sensitive.OutPut
	Msg.Msg = openAiRcv.Choices[0].Message.Content
	out.PutOutMsg(&Msg)
	time.Sleep(20 * time.Second)
	return
}

// 使用函数，进行第二次调用
func secondRequest(firstRequest postDataWithFunction, firstResp OpenAiRcv) OpenAiRcv {
	funcName := firstResp.Choices[0].Message.FunctionCall.Name
	// 定义一个正则表达式，用于匹配每一对双引号中间的内容
	pattern := `"(.*?)"`

	// 使用正则表达式解析 JSON 字符串中的字段值
	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch(firstResp.Choices[0].Message.FunctionCall.Arguments, -1)

	// 遍历字符串切片，并取出所有偶数位置上的元素
	var values []string
	for i, match := range matches {
		if i%2 == 1 { // 只保留偶数位置上的元素
			values = append(values, match[1])
		}
	}
	if values == nil {
		log.Println("函数调用内容解析失败")
		return OpenAiRcv{}
	}
	result := function.GetFunctionResult(funcName, values)
	ms := RequestMessages{
		Role:    "function",
		Name:    funcName,
		Content: result,
	}
	firstRequest.Messages = append(firstRequest.Messages, ms)
	tempRequest := postDataWithoutFunction{
		Model:            firstRequest.Model,
		Messages:         firstRequest.Messages,
		MaxTokens:        firstRequest.MaxTokens,
		Temperature:      firstRequest.Temperature,
		TopP:             firstRequest.TopP,
		Stop:             firstRequest.Stop,
		PresencePenalty:  firstRequest.PresencePenalty,
		FrequencyPenalty: firstRequest.FrequencyPenalty,
		User:             firstRequest.User,
	}
	postDataBytes, err := json.Marshal(tempRequest)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
		return OpenAiRcv{}
	}
	req, _ := http.NewRequest("POST", OpenAIChatUrl, bytes.NewBuffer(postDataBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GPTCfg.OpenAi.ApiKey)
	client, err := proxy.Client()
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
		return OpenAiRcv{}
	}
	defer resp.Body.Close()
	if resp == nil {
		backend.WebsocketToNLP <- true
		log.Println("response is nil")
		return OpenAiRcv{}
	}
	body, _ := io.ReadAll(resp.Body)
	var openAiRcv OpenAiRcv
	err = json.Unmarshal(body, &openAiRcv)
	if err != nil {
		log.Println(err)
		return OpenAiRcv{}
	}
	if len(openAiRcv.Choices) == 0 {
		backend.WebsocketToNLP <- true
		log.Println("OpenAI API调用失败，返回内容：", string(body))
		return OpenAiRcv{}
	}
	return openAiRcv
}
