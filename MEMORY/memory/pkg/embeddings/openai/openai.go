package oai

import (
	"GoTuber/NLP/config"
	"GoTuber/proxy"
	"bytes"
	"context"
	"encoding/json"
	"github.com/aldarisbm/memory/pkg/embeddings"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
)

const openaiApiUrl = "https://api.openai.com/v1/embeddings"

type Embedding struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int
	TotalTokens      int `json:"total_tokens"`
}

type EmbeddingResponse struct {
	Object string      `json:"object"`
	Data   []Embedding `json:"data"`
	Model  string      `json:"model"`
	Usage  Usage       `json:"usage"`
}

type embedder struct {
	c     *openai.Client
	model openai.EmbeddingModel
}

// NewOpenAIEmbedder returns an Embedder that uses OpenAI's API to embed text.
func NewOpenAIEmbedder(opts ...CallOptions) *embedder {
	o := applyCallOptions(opts, options{
		model: openai.AdaEmbeddingV2,
	})
	c := openai.NewClient(o.apiKey)
	return &embedder{
		c:     c,
		model: o.model,
	}
}

// EmbedDocumentText embeds the given text
func (e *embedder) EmbedDocumentText(text string) ([]float32, error) {
	postData := openai.EmbeddingRequest{
		Input: []string{text},
		Model: e.model,
	}
	//resp, err := e.c.CreateEmbeddings(ctx, req)
	postDataBytes, err := json.Marshal(postData)
	req, _ := http.NewRequest("POST", openaiApiUrl, bytes.NewBuffer(postDataBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.GPTCfg.OpenAi.ApiKey)
	client, err := proxy.Client()
	if err != nil {
		log.Println("embed处理错误，", err)
	}
	rep, err := client.Do(req)
	if err != nil {
		log.Println("embed处理错误，", err)
		return nil, err
	}
	if err != nil {
		log.Println("embed处理错误，", err)
		return nil, err
	}
	defer rep.Body.Close()
	if rep == nil {
		log.Println("response is nil")
		return nil, err
	}
	body, _ := io.ReadAll(rep.Body)
	var resp EmbeddingResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("embed处理错误，", err)
		return nil, err
	}
	return resp.Data[0].Embedding, nil
}

// EmbedDocumentTexts embeds the given texts
func (e *embedder) EmbedDocumentTexts(texts []string) ([][]float32, error) {
	ctx := context.Background()
	req := openai.EmbeddingRequest{
		Input: make([]string, len(texts)),
		Model: e.model,
	}
	for i, text := range texts {
		req.Input[i] = text
	}
	resp, err := e.c.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, err
	}
	embeddings := make([][]float32, len(resp.Data))
	for i, data := range resp.Data {
		embeddings[i] = data.Embedding
	}
	return embeddings, nil
}

// Ensure embedder implements embeddings.Embedder
var _ embeddings.Embedder = (*embedder)(nil)
