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
	input := pinecone.Input{
		Type:      chat.Type,
		Namespace: chat.Namespace,
		UserName:  chat.UserName,
		Human:     chat.Human,
		AI:        chat.AI,
		Vector:    vector,
	}
	input.PineconeStore()
}

func (chat Chat) GetMemory() (humanText, aiText, user string) {
	vector := embedding.OpenaiEmbedding(chat.Human)
	if vector == nil {
		return "", "", ""
	}
	memory := pinecone.PineconeQuery("chat", "live", vector) //什么默认字段
	if memory == nil {
		return "", "", ""
	}
	return memory[1], memory[0], memory[3] //索引0：AI的回答;1：用户提问;2:类型;3:用户名
}
