package MEMORY

import (
	"GoTuber/MEMORY/memory/vector_database/pinecone"
)

func InitMemory() {
	InitMemoryConfig()
	if MemoryCfg.IsUse {
		pinecone.InitPineconeConfig()
	}
}
