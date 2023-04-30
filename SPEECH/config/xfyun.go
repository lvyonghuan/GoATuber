package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Xfyun struct {
	Xfyun struct {
	}
}

var XFCfg Xfyun

func InitXFConfig() {
	if _, err := os.Stat("SPEECH/config/XFConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("SPEECH/config/XFConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# 语音模块通用设置\n[speech]\n" +
			"# 使用科大讯飞语音合成平台（调用在线接口）（目前默认）\n" +
			"use_xfyun = true \n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("XFConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./SPEECH/config") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&XFCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
