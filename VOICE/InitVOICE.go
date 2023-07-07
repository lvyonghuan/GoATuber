package VOICE

import (
	"GoTuber/VOICE/config"
	"GoTuber/VOICE/service/azure"
)

// InitVOICE 初始化语音模块
func InitVOICE() {
	config.InitVOICEConfig()
	if config.VoiceCfg.Voice.UseAzure {
		go azure.GetAuthentication()
	}
}
