package embedding

import (
	"GoTuber/NLP/config"
	"GoTuber/proxy"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RepsFromOpenai struct {
	Object string `json:"object"`
	Data   []Data `json:"data"`
	Model  string `json:"model"`
	Usage  Usage  `json:"usage"`
}
type Data struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}
type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type req struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// OpenaiEmbedding 调用openai生成embedding
func OpenaiEmbedding(msg string) []float32 {
	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}
	var request req
	request.Input = msg
	request.Model = "text-embedding-ada-002"
	data, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/embeddings", bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return nil
	}
	req.Header.Set("Authorization", "Bearer "+config.GPTCfg.OpenAi.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("embedding生成错误:", err)
		return nil
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("embedding生成错误:", err)
		return nil
	}
	var reps RepsFromOpenai
	err = json.Unmarshal(bodyText, &reps)
	if err != nil || len(reps.Data) == 0 {
		log.Println("embedding生成错误:", err)
		return nil
	}
	return reps.Data[0].Embedding
}

// AzureOpenaiEmbedding 调用azure生成embedding
func AzureOpenaiEmbedding(msg string) []float32 {
	url := "https://" + config.GPTCfg.Azure.EndPoint + "openai/deployments/" + config.GPTCfg.Azure.Memory.DeploymentId + "/embeddings?api-version=" + config.GPTCfg.Azure.Memory.ApiVersion
	type req struct {
		Input string `json:"input"`
	}
	var request req
	request.Input = msg
	data, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return nil
	}

	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}

	//设置请求头
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("api-key", config.GPTCfg.Azure.ApiKey)

	resp, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	var reps RepsFromOpenai
	err = json.Unmarshal(bodyText, &reps)
	if err != nil || len(reps.Data) == 0 {
		log.Println(err)
		return nil
	}
	return reps.Data[0].Embedding
}
