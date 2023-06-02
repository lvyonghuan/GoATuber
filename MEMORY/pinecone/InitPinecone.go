package pinecone

import (
	memory "GoTuber/MEMORY/memory/pkg"
	oai "GoTuber/MEMORY/memory/pkg/embeddings/openai"
	pc "GoTuber/MEMORY/memory/pkg/vectorstore/pinecone"
	"GoTuber/NLP/config"
)

type MemoryData struct {
	ID         string                 `json:"ID"`
	User       string                 `json:"User"`
	Text       string                 `json:"Text"`
	CreatedAt  string                 `json:"CreatedAt"`
	LastReadAt string                 `json:"LastReadAt"`
	Vector     []float64              `json:"Vector"`
	Metadata   map[string]interface{} `json:"Metadata"`
}

var Memory *memory.Memory

func InitPinecone() {
	vs := pc.NewStorer(
		pc.WithApiKey(pineconeCfg.ApiKey),
		pc.WithIndexName(pineconeCfg.IndexName),
		pc.WithProjectName(pineconeCfg.ProjectName),
		pc.WithEnvironment(pineconeCfg.Environment),
	)
	openAI := oai.NewOpenAIEmbedder(
		oai.WithApiKey(config.GPTCfg.OpenAi.ApiKey),
	)
	Memory = memory.NewMemory(memory.WithVectorStore(vs), memory.WithEmbedder(openAI))
}
