package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type VOICE struct {
	Voice struct {
		UseXfyun        bool `mapstructure:"use_xfyun"`
		UseTalkinggenie bool `mapstructure:"use_talkinggenie"`
		UseAzure        bool `mapstructure:"use_azure"`
	}
}

var VoiceCfg VOICE

// InitVOICEConfig 初始化语音模块配置文件
func InitVOICEConfig() {
	if _, err := os.Stat("config/VOICE/VoiceConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/VOICE/VoiceConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# 语音模块通用设置\n[voice]\n" +
			"# 使用科大讯飞语音合成平台（调用在线接口）\n" +
			"use_xfyun = true \n" +
			"# 使用会话精灵（非官方api）（调用在线接口）\n" +
			"use_talkinggenie = false \n" +
			"# 使用azure\n" +
			"use_azure = false\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("VoiceConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/VOICE") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&VoiceCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
	if VoiceCfg.Voice.UseXfyun {
		InitXFConfig()
	} else if VoiceCfg.Voice.UseTalkinggenie {
		InitTalkinggenieConfig()
	} else if VoiceCfg.Voice.UseAzure {
		InitAzureConfig()
	}
}
