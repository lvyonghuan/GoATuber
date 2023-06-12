package memory_gpt

import (
	"GoTuber/MEMORY/memory/embedding"
	"GoTuber/MEMORY/memory/vector_database/pinecone"
)

type Chat struct {
	Type      string
	Namespace string
	UserName  string
	Human     string
	AI        string
}

//TODO:优化记忆逻辑

func (chat Chat) StoreMessage() {
	vector := embedding.OpenaiEmbedding(chat.Human)
	if vector == nil {
		return
	}
	pinecone.PineconeStore(chat.Type, chat.Human, chat.UserName, chat.Namespace, vector)
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
	return memory[2], memory[0] //TODO:这里有点太恶俗了，要重新理一理
}
