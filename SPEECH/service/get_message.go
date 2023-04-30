package service

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/SPEECH/config"
	"GoTuber/SPEECH/service/talkinggenie"
)

func GetMessage(msg sensitive.OutPut) {
	if config.SpeechCfg.Speech.UseXfyun {
		//TODO：todo
	} else if config.SpeechCfg.Speech.UseTalkinggenie {
		talkinggenie.GetVoice(msg)
	}
}
