package SPEECH

import (
	"GoTuber/SPEECH/config"
	"GoTuber/SPEECH/service/azure"
)

// InitSPEECH 初始化语音模块
func InitSPEECH() {
	config.InitSPEECHConfig()
	if config.SpeechCfg.Speech.UseAzure {
		go azure.GetAuthentication()
	}
}
