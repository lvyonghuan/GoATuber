package talkinggenie

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/VOICE/config"
	"strconv"
)

func GetVoice(Msg *sensitive.OutPut) {
	msg := Msg.Msg
	speed := strconv.FormatFloat(config.TalkinggenieCfg.Talkinggenie.Speed, 'f', 2, 64)
	volume := strconv.FormatFloat(config.TalkinggenieCfg.Talkinggenie.Volume, 'f', 2, 64)
	url := "https://dds.dui.ai/runtime/v1/synthesize?voiceId=" + config.TalkinggenieCfg.Talkinggenie.VoiceId + "&text=" + msg + "&speed=" + speed + "&volume=" + volume + "&audioType=wav"
	Msg.Mu.Lock()
	Msg.Voice = url
	Msg.VType = 1
	Msg.Mu.Unlock()
}
