package service

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/VOICE/config"
	"GoTuber/VOICE/service/azure"
	"GoTuber/VOICE/service/talkinggenie"
	"GoTuber/VOICE/service/xfyun"
)

func GetMessage(msg *sensitive.OutPut) {
	if config.VoiceCfg.Voice.UseXfyun {
		xfyun.GetVoice(msg)
	} else if config.VoiceCfg.Voice.UseTalkinggenie {
		talkinggenie.GetVoice(msg)
	} else if config.VoiceCfg.Voice.UseAzure {
		azure.GetVoice(msg)
	}
}
