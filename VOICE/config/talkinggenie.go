package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Talkinggenie struct {
	Talkinggenie struct {
		Speed   float64 `mapstructure:"speed"`   //语速，值和语速成正比
		Volume  float64 `mapstructure:"volume"`  //声音大小
		VoiceId string  `mapstructure:"voiceId"` //发音类型，参考https://czyt.tech/post/a-free-tts-api/
	}
}

var TalkinggenieCfg Talkinggenie

func InitTalkinggenieConfig() {
	if _, err := os.Stat("config/VOICE/talkinggenieConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/VOICE/talkinggenieConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# talkinggenie参数设置\n[talkinggenie]\n" +
			"# 语速 \n" +
			"speed = 1\n" +
			"# 声量\n" +
			"volume = 50\n" +
			"# 发音类型\n" +
			"voiceId = \"qiumum_0gushi\"\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("talkinggenieConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/VOICE") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&TalkinggenieCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
