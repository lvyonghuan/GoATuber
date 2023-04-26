package NLP

import "GoTuber/NLP/service"

func InitNLP() {
	service.HandelMsg.IsUse = false
	go service.ChooseMessage()
	service.ReadToGetFlag <- true
	go service.Read()
}
