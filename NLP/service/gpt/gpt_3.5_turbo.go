package gpt

import (
	"GoTuber/MEMORY"
	memory_gpt "GoTuber/MEMORY/LLMmodel/gpt"
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/NLP/service/out"
	backend "GoTuber/frontend/live2d_backend"
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/config"
	"GoTuber/proxy"
)

var MS []Messages     //向OpenAI传递的消息，包含了用户设定的提示词
var roleMS []Messages //角色信息

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

// InitRole 这个函数用于获取role.cfg的角色文本信息
func InitRole() {
	file, err := os.Open("./config/NLP/GPTConfig/role.cfg")
	if err != nil {
		log.Fatalf("open config file failed: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for i := 1; ; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if i >= 2 {
			msg := strings.Split(line, ":")
			ms := &Messages{
				Role:    msg[0],
				Content: msg[1],
			}
			roleMS = append(roleMS, *ms)
		}
	}
	MS = append(MS, roleMS...)
}

// GenerateText 文本请求
func GenerateText(msg *model.Msg) {
	log.Println("正在生成文本......")

	//记忆相关
	memory := memory_gpt.Chat{
		Human: msg.Msg,
		AI:    "",
	}
	if MEMORY.MemoryCfg.IsUse {
		user, text := memory.GetMemory()
		mem := Messages{
			Role:    "system",
			Content: "你是一个虚拟主播。你可以选择利用这些记忆，当记忆无关的时候，也可以选择忽略。以下是记忆部分。" + user + "说，" + text,
		}
		MS = append(MS, mem)
	}

	messages := &Messages{
		Role:    "user",
		Content: msg.Name + "说：" + msg.Msg,
	}
	MS = append(MS, *messages)
	postDataTemp := postData{
		Model:            config.GPTCfg.OpenAi.Model,
		Messages:         MS,
		MaxTokens:        config.GPTCfg.OpenAi.MaxTokens,
		Temperature:      config.GPTCfg.OpenAi.Temperature,
		TopP:             config.GPTCfg.OpenAi.TopP,
		Stop:             config.GPTCfg.OpenAi.Stop,
		PresencePenalty:  config.GPTCfg.OpenAi.PresencePenalty,
		FrequencyPenalty: config.GPTCfg.OpenAi.FrequencyPenalty,
	}
	postDataBytes, err := json.Marshal(postDataTemp)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println(err)
		return
	}
	req, _ := http.NewRequest("POST", Openaiapiurl1, bytes.NewBuffer(postDataBytes))
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
	openAiRcv.Choices[0].Message.Content = strings.Replace(openAiRcv.Choices[0].Message.Content, "\n\n", "", 1)
	log.Printf("Model: %s TotalTokens: %d+%d=%d", openAiRcv.Model, openAiRcv.Usage.PromptTokens, openAiRcv.Usage.CompletionTokes, openAiRcv.Usage.TotalTokens)
	//TODO：保留了短期记忆，不过消耗的token超过一个阈值的时候会执行删除。计划由用户设定这个功能。也许可以加入一个比较连续的短期记忆功能。
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
	time.Sleep(20 * time.Second)
	return
}
