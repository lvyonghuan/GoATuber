package SPEECH

import (
	"GoTuber/proxy"
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//model部分

// config
type azureConfig struct {
	Key          string `mapstructure:"key"`
	SpeechRegion string `mapstructure:"speech_region"`
	Language     string `mapstructure:"language"`
	Format       string `mapstructure:"format"`
}

type response struct {
	RecognitionStatus string `json:"RecognitionStatus"`
	Offset            int    `json:"Offset"`
	Duration          int    `json:"Duration"`
	DisplayText       string `json:"DisplayText"`
}

var AzureCfg azureConfig

// config部分

func initAzureSpeechConfig() {
	if _, err := os.Stat("config/SPEECH/azure/azure_config.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/SPEECH/azure/azure_config.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"key = \"xxxxxx\" \n" +
			"# 参考https://learn.microsoft.com/zh-cn/azure/cognitive-services/speech-service/regions" +
			"speech_region = \"eastus\" \n" +
			"language = \"zh-CN\"" +
			"format = \"simple \""))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("azure_config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/SPEECH/azure") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&AzureCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}

//service部分

func voiceToTextByAzure(voice *bytes.Buffer) (msg string) {
	url := "https://" + AzureCfg.SpeechRegion + ".stt.speech.microsoft.com/speech/recognition/conversation/cognitiveservices/v1?language=" + AzureCfg.Language + "&format=" + AzureCfg.Format
	req, err := http.NewRequest("POST", url, voice)
	if err != nil {
		log.Println("创建request错误:", err)
		return ""
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", AzureCfg.Key)
	req.Header.Set("Content-Type", "audio/wav")
	client, err := proxy.Client()
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("请求错误：", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	var response response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("json反序列化失败：", err)
		return ""
	}
	if response.RecognitionStatus != "Success" {
		log.Println("azure语音识别失败，错误信息：", response.RecognitionStatus)
	}
	log.Println(response.DisplayText)
	msg = response.DisplayText
	return msg
}
