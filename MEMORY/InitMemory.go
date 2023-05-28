package MEMORY

import "GoTuber/MEMORY/pinecone"

func InitMemory() {
	InitMemoryConfig()
	if MemoryCfg.IsUse {
		pinecone.InitPineconeConfig()
		pinecone.InitPinecone()
	}
}
