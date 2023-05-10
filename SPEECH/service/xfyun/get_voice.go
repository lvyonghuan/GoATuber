package xfyun

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/SPEECH/config"
	"log"
)

const status = 2 //数据状态，固定为2

type voice struct {
	common struct {
		AppID string `json:"app_id"` //用户的APP_ID
	}
	business struct {
		Aue    string `json:"aue"`    //音频编码
		Sft    int    `json:"sft"`    //流式返回
		Auf    string `json:"auf"`    //音频采样率
		Vcn    string `json:"vcn"`    //发音人
		Speed  int    `json:"speed"`  //语速
		Volume int    `json:"volume"` //音量
		Pitch  int    `json:"pitch"`  //音高
		Bgs    int    `json:"bgs"`    //是否有背景音，肯定没有啊，，，
		//Tte    string `json:"tte"`    //文件编码格式
		Reg string `json:"reg"` //英文发音方式
		Rdn string `json:"rdn"` //数字发音方式
	}
	data struct {
		Text   string `json:"text"`   //文本内容
		Status int    `json:"status"` //状态
	}
}

func GetVoice(text *sensitive.OutPut) {
	conn := connectXFYun()
	defer conn.Close()
	if conn == nil {
		return
	}
	var v voice
	v.common.AppID = config.XFCfg.Xfyun.ApiKey
	v.business.Aue = config.XFCfg.XfyunVoice.Aue
	v.business.Sft = config.XFCfg.XfyunVoice.Sft
	v.business.Auf = config.XFCfg.XfyunVoice.Auf
	v.business.Vcn = config.XFCfg.XfyunVoice.Vcn
	v.business.Speed = config.XFCfg.XfyunVoice.Speed
	v.business.Volume = config.XFCfg.XfyunVoice.Volume
	v.business.Pitch = config.XFCfg.XfyunVoice.Pitch
	v.business.Bgs = 0
	v.business.Reg = config.XFCfg.XfyunVoice.Reg
	v.business.Rdn = config.XFCfg.XfyunVoice.Rdn
	v.data.Text = text.Msg
	v.data.Status = 2
	err := conn.WriteJSON(v)
	if err != nil {
		log.Println("向讯飞发送信息错误，错误信息：", err)
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("读取讯飞语音接口返回消息错误，错误信息：%v", err)
		}
		log.Println(msg)
		return
	}
}
