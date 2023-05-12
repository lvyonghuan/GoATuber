package xfyun

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/SPEECH/config"
	"encoding/base64"
	"encoding/json"
	"log"
)

const status = 2 //数据状态，固定为2

type resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Audio  string `json:"audio,omitempty"`
		Status int    `json:"status,omitempty"`
		Ced    string `json:"ced,omitempty"`
	} `json:"data"`
	Sid string `json:"sid"`
}

func GetVoice(text *sensitive.OutPut) {
	conn := connectXFYun()
	defer conn.Close()
	if conn == nil {
		return
	}
	request := map[string]interface{}{
		"common": map[string]interface{}{
			"app_id": config.XFCfg.Xfyun.AppID, //用户的APP_ID
		},
		"business": map[string]interface{}{
			"aue":    config.XFCfg.XfyunVoice.Aue,    //音频编码
			"sfl":    config.XFCfg.XfyunVoice.Sfl,    //流式返回
			"auf":    config.XFCfg.XfyunVoice.Auf,    //音频采样率
			"vcn":    config.XFCfg.XfyunVoice.Vcn,    //发音人
			"speed":  config.XFCfg.XfyunVoice.Speed,  //语速
			"volume": config.XFCfg.XfyunVoice.Volume, //音量
			"pitch":  config.XFCfg.XfyunVoice.Pitch,  //音高
			"bgs":    0,                              //是否有背景音，肯定没有啊，，，
			"reg":    config.XFCfg.XfyunVoice.Reg,    //英文发音方式
			"rdn":    config.XFCfg.XfyunVoice.Rdn,    //数字发音方式
		},
		"data": map[string]interface{}{
			"text":     base64.StdEncoding.EncodeToString([]byte(text.Msg)), //文本内容
			"encoding": "UTF8",
			"status":   status, //状态
		},
	}
	err := conn.WriteJSON(request)
	if err != nil {
		log.Println("向讯飞发送信息错误，错误信息：", err)
		return
	}
	var voice string
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("读取讯飞语音接口返回消息错误，错误信息：%v", err)
		}
		var respMsg resp
		err = json.Unmarshal(msg, &respMsg)
		if err != nil {
			log.Fatalf("讯飞返回信息格式化失败，错误信息：%v", err)
		}
		voice += respMsg.Data.Audio
		if respMsg.Data.Status == 2 {
			break
		}
	}
	text.Mu.Lock()
	text.Voice = voice
	text.VType = 2
	text.Mu.Unlock()
}
