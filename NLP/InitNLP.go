package NLP

import (
	"GoTuber/NLP/config"
	"GoTuber/NLP/service"
	"GoTuber/NLP/service/gpt"
	"GoTuber/frontend/live2d_backend"
)

func InitNLP() {
	service.HandelMsg.IsUse = false
	go service.ChooseMessage()
	service.ReadToGetFlag <- true
	backend.WebsocketToNLP <- true //初始化NLP信息读取模块
	go service.HandelMessage()
	config.InitNLPConfig()
	if config.NLPCfg.Nlp.UseGPT == true {
		config.InitGPTConfig()
		gpt.InitRole()
	}
}
