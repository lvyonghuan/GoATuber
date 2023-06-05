package gpt

import (
	"GoTuber/MEMORY"
	memory_gpt "GoTuber/MEMORY/NLPmodel/gpt"
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/config"
	"GoTuber/NLP/service/out"
	backend "GoTuber/frontend/live2d_backend"
	"GoTuber/proxy"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func GenerateTextByAzureOpenAI(msg *model.Msg) {
	url := config.GPTCfg.Azure.EndPoint + "openai/deployments/" + config.GPTCfg.Azure.DeploymentID + "/chat/completions?api-version=" + config.GPTCfg.Azure.ApiVersion
	//记忆相关
	memory := memory_gpt.Chat{
		Human: msg.Msg,
		AI:    "",
	}
	if MEMORY.MemoryCfg.IsUse {
		user, text := memory.GetMemory()
		mem := Messages{
			Role:    "system",
			Content: "你是一个虚拟主播。你可以选择利用这些记忆，当记忆无关的时候，也可以选择忽略。请不要在发言中直接提到“记忆”。以下是记忆部分。" + user + "说，" + text,
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
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(postDataBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", config.GPTCfg.Azure.ApiKey)
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
		log.Println("Azure OpenAI API调用失败，返回内容：", string(body))
		return
	}
	openAiRcv.Choices[0].Message.Content = strings.Replace(openAiRcv.Choices[0].Message.Content, "\n\n", "", 1)
	log.Printf("Model: %s TotalTokens: %d+%d=%d", openAiRcv.Model, openAiRcv.Usage.PromptTokens, openAiRcv.Usage.CompletionTokes, openAiRcv.Usage.TotalTokens)
	if openAiRcv.Usage.TotalTokens > 500 {
		MS = MS[:0]
		MS = append(MS, roleMS...)
	}

	if MEMORY.MemoryCfg.IsUse {
		memory.UserName = msg.Name
		memory.AI = openAiRcv.Choices[0].Message.Content
		go memory.StoreMessage()
	}

	var Msg sensitive.OutPut
	Msg.Msg = openAiRcv.Choices[0].Message.Content
	out.PutOutMsg(&Msg)
	return
}
