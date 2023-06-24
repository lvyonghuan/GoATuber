package NLP

import (
	"GoTuber/NLP/config"
	"GoTuber/NLP/service"
	"GoTuber/NLP/service/gpt"
	"GoTuber/NLP/service/gpt/function"
)

func InitNLP() {
	service.HandelMsg.IsUse = false
	service.MsgMu.Lock() //初始化锁
	go service.ChooseMessage()
	service.ReadToGetFlag <- true
	go service.HandelMessage()
	config.InitNLPConfig()
	if config.NLPCfg.Nlp.UseGPT == true || config.NLPCfg.Nlp.UseAzureGPT == true {
		function.InitFunction()
		function.InitFunctionJson()
		config.InitGPTConfig()
		gpt.InitRole()
	}
	if config.NLPCfg.Nlp.UseLocalModel {
		config.InitNLPLocalConfig()
	}
}
