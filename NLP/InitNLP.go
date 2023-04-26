package NLP

import "GoTuber/NLP/service"

func InitNLP() {
	service.HandelMsg.IsUse = false
	go service.ChooseMessage()
	go service.Read()
}
