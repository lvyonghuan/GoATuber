package SPEECH

import (
	"GoTuber/MESSAGE/model"
	"bytes"
	"log"
	"os"
)

func GetVoice(msg *model.Msg) {
	audioFilePath := "./dist/temp/record.wav"
	file, err := os.ReadFile(audioFilePath)
	if err != nil {
		log.Println("读取录音文件错误：", err)
		return
	}
	fileData := bytes.NewBuffer(file)
	message := handelVoice(fileData)
	if message == "" {
		return
	}
	msg.Msg = message
	msg.Uid = "0"
	msg.Name = "speaker"
}

func handelVoice(fileData *bytes.Buffer) (msg string) {
	if SpeechCfg.UseAzure {
		msg = voiceToTextByAzure(fileData)
	} else if SpeechCfg.UseOther {

	}
	return msg
}
