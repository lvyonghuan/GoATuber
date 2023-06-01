package memory_gpt

import (
	"GoTuber/MEMORY/pinecone"
	"log"
)

type Chat struct {
	UserName string
	Human    string
	AI       string
}

//TODO:优化记忆逻辑

func (chat Chat) StoreMessage() {
	doc := pinecone.Memory.NewDocument(chat.Human, chat.UserName)
	if err := pinecone.Memory.StoreDocument(doc); err != nil {
		log.Println("存储用户消息失败，错误信息：", err)
		return
	}
	doc = pinecone.Memory.NewDocument(chat.AI, "你")
	if err := pinecone.Memory.StoreDocument(doc); err != nil {
		log.Println("存储AI消息失败，错误信息：", err)
		return
	}
}

func (chat Chat) GetMemory() (user, text string) {
	docs, err := pinecone.Memory.RetrieveSimilarDocumentsByText(chat.Human, 1)
	if err != nil {
		log.Println("记忆获取失败，错误消息：", err)
	}
	for _, d := range docs {
		log.Println(d.User, d.Text)
		return d.User, d.Text
	}
	return "", ""
}
