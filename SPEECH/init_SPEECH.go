package SPEECH

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// model
type speechConfig struct {
	UseAzure bool `mapstructure:"use_azure"` //使用azure
	UseOther bool `mapstructure:"use_other"` //使用其它
}

var SpeechCfg speechConfig

//config

func InitConfig() {
	if _, err := os.Stat("./config/SPEECH/speech_config.cfg"); os.IsNotExist(err) {
		f, err := os.Create("./config/SPEECH/speech_config.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"use_azure = true \n" +
			"use_other = false"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("speech_config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/SPEECH") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&SpeechCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
	if SpeechCfg.UseAzure {
		initAzureSpeechConfig()
	} else if SpeechCfg.UseOther {
		//TODO：
	}
}
