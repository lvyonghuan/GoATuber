package gpt

import (
	memory_gpt "GoTuber/MEMORY/NLPmodel/gpt"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"GoTuber/MEMORY"
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/config"
	"GoTuber/NLP/service/out"
	backend "GoTuber/frontend/live2d_backend"
	"GoTuber/proxy"
)

// GenerateTextByOpenAI 文本请求
func GenerateTextByOpenAI(msg *model.Msg) {
	//记忆相关
	memory := memory_gpt.Chat{
		Human: msg.Msg,
		AI:    "",
	}
	if MEMORY.MemoryCfg.IsUse {
		user, text := memory.GetMemory()
		mem := Messages{
			Role:    "system",
			Content: user + "说，" + text,
		}
		MS = append(MS, mem)
	}

	messages := &Messages{
		Role:    "user",
		Content: msg.Name + "说：" + msg.Msg,
	}
	MS = append(MS, *messages)
	postDataTemp := postData{
		Model:            config.GPTCfg.General.Model,
		Messages:         MS,
		MaxTokens:        config.GPTCfg.General.MaxTokens,
		Temperature:      config.GPTCfg.General.Temperature,
		TopP:             config.GPTCfg.General.TopP,
		Stop:             config.GPTCfg.General.Stop,
		PresencePenalty:  config.GPTCfg.General.PresencePenalty,
		FrequencyPenalty: config.GPTCfg.General.FrequencyPenalty,
		User:             msg.Name,
	}
	postDataBytes, err := json.Marshal(postDataTemp)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
		return
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
		return
	}
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
		log.Println(err)
		return
	}
	if len(openAiRcv.Choices) == 0 {
		backend.WebsocketToNLP <- true
		log.Println("OpenAI API调用失败，返回内容：", string(body))
		return
	}
	log.Println("openai生成内容：", openAiRcv.Choices[0].Message.Content)
	openAiRcv.Choices[0].Message.Content = strings.Replace(openAiRcv.Choices[0].Message.Content, "\n\n", "", 1)
	log.Printf("Model: %s TotalTokens: %d+%d=%d", openAiRcv.Model, openAiRcv.Usage.PromptTokens, openAiRcv.Usage.CompletionTokes, openAiRcv.Usage.TotalTokens)

	//压入AI的回答，形成短期记忆
	messagesAI := Messages{
		Role:    "assistant",
		Content: openAiRcv.Choices[0].Message.Content,
	}
	MS = append(MS, messagesAI)

	if MEMORY.MemoryCfg.IsUse {
		memory.UserName = msg.Name
		memory.Type = "chat"
		memory.Namespace = "live"
		memory.AI = openAiRcv.Choices[0].Message.Content
		go memory.StoreMessage()
		cleanMemoryMessage() //清除这一次对话的记忆内容
	}

	//TODO：保留了短期记忆，不过消耗的token超过一个阈值的时候会执行删除。计划由用户设定这个功能。也许可以加入一个比较连续的短期记忆功能。
	if openAiRcv.Usage.TotalTokens > 500 {
		cleanAllMessage()
	}

	var Msg sensitive.OutPut
	Msg.Msg = openAiRcv.Choices[0].Message.Content
	out.PutOutMsg(&Msg)
	time.Sleep(20 * time.Second)
	return
}
