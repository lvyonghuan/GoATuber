package gpt

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MOOD"
	"GoTuber/SPEECH/service"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/config"
	"GoTuber/proxy"
)

const Openaiapiurl1 = "https://api.openai.com/v1/chat/completions" //对话使用的url

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 对话使用的Request body
type postData struct {
	Model            string     `json:"model"`
	Messages         []Messages `json:"messages"` //message依靠传入信息获取
	MaxTokens        int        `json:"max_tokens"`
	Temperature      float64    `json:"temperature"`
	TopP             float64    `json:"top_p"`
	Stop             string     `json:"stop"`
	PresencePenalty  float64    `json:"presence_penalty"`
	FrequencyPenalty float64    `json:"frequency_penalty"`
}

// OpenAiRcv 对话使用的Response
type OpenAiRcv struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message      Messages `json:"message"`
		FinishReason string   `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens    int `json:"prompt_tokens"`
		CompletionTokes int `json:"completion_tokens"`
		TotalTokens     int `json:"total_tokens"`
	}
}

// GenerateText 文本请求
func GenerateText(msg *model.Msg) {
	log.Println("正在生成文本......")
	var ms []Messages
	messages := &Messages{
		Role:    "user",
		Content: msg.Msg,
	}
	ms = append(ms, *messages)
	postDataTemp := postData{
		Model:            config.GPTCfg.OpenAi.Model,
		Messages:         ms,
		MaxTokens:        config.GPTCfg.OpenAi.MaxTokens,
		Temperature:      config.GPTCfg.OpenAi.Temperature,
		TopP:             config.GPTCfg.OpenAi.TopP,
		Stop:             config.GPTCfg.OpenAi.Stop,
		PresencePenalty:  config.GPTCfg.OpenAi.PresencePenalty,
		FrequencyPenalty: config.GPTCfg.OpenAi.FrequencyPenalty,
	}
	postDataBytes, err := json.Marshal(postDataTemp)
	if err != nil {
		log.Println(err)
		return
	}
	req, _ := http.NewRequest("POST", Openaiapiurl1, bytes.NewBuffer(postDataBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GPTCfg.OpenAi.ApiKey)
	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp == nil {
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
		log.Println("OpenAI API调用失败，返回内容：", string(body))
		return
	}
	openAiRcv.Choices[0].Message.Content = strings.Replace(openAiRcv.Choices[0].Message.Content, "\n\n", "\n", 1)
	log.Printf("Model: %s TotalTokens: %d+%d=%d", openAiRcv.Model, openAiRcv.Usage.PromptTokens, openAiRcv.Usage.CompletionTokes, openAiRcv.Usage.TotalTokens)
	var Msg sensitive.OutPut
	Msg.Msg = openAiRcv.Choices[0].Message.Content
	Msg.AIFilter(&Msg)
	go MOOD.GetMessage(Msg)
	service.GetMessage(Msg)
	time.Sleep(20 * time.Second)
	return
}
