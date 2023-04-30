package talkinggenie

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/SPEECH/config"
	"log"
	"strconv"
)

func GetVoice(Msg sensitive.OutPut) {
	msg := Msg.Msg
	url := "https://dds.dui.ai/runtime/v1/synthesize?voiceId=" + config.TalkinggenieCfg.Talkinggenie.VoiceId + "&text=" + msg + "&speed=" + strconv.Itoa(config.TalkinggenieCfg.Talkinggenie.Speed) + "&volume=" + strconv.Itoa(config.TalkinggenieCfg.Talkinggenie.Volume) + "&audioType=wav"
	log.Println(url)
}
