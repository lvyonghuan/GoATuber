package NLP

import (
	"GoTuber/NLP/config"
	"GoTuber/NLP/service"
)

func InitNLP() {
	service.HandelMsg.IsUse = false
	go service.ChooseMessage()
	service.ReadToGetFlag <- true
	go service.HandelMessage()
	config.InitNLPConfig()
	config.InitGPTConfig()
}
