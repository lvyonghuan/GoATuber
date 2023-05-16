package service

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/SPEECH/config"
	"GoTuber/SPEECH/service/azure"
	"GoTuber/SPEECH/service/talkinggenie"
	"GoTuber/SPEECH/service/xfyun"
)

func GetMessage(msg *sensitive.OutPut) {
	if config.SpeechCfg.Speech.UseXfyun {
		xfyun.GetVoice(msg)
	} else if config.SpeechCfg.Speech.UseTalkinggenie {
		talkinggenie.GetVoice(msg)
	} else if config.SpeechCfg.Speech.UseAzure {
		azure.GetVoice(msg)
	}
}
