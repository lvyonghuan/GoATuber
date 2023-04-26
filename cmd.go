package main

import (
	"GoTuber/CHAT"
	"GoTuber/MESSAGE"
	"GoTuber/NLP"
)

func main() {
	go MESSAGE.GetMessage()
	NLP.InitNLP()
	CHAT.InitChat()
}
