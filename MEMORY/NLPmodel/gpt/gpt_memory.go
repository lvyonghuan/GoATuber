package memory_gpt

import (
	"GoTuber/MEMORY/memory/embedding"
	"GoTuber/MEMORY/memory/vector_database/pinecone"
)

type Chat struct {
	UserName string
	Human    string
	AI       string
}

//TODO:优化记忆逻辑

func (chat Chat) StoreMessage() {
	vector := embedding.OpenaiEmbedding(chat.Human)
	if vector == nil {
		return
	}
	pinecone.PineconeStore("chat", chat.Human, chat.UserName, "live", vector)
}

func (chat Chat) GetMemory() (user, text string) {
	vector := embedding.OpenaiEmbedding(chat.Human)
	if vector == nil {
		return "", ""
	}
	memory := pinecone.PineconeQuery("chat", "live", vector)
	if memory == nil {
		return "", ""
	}
	return memory[2], memory[0]
}
