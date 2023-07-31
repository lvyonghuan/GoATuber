package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type NLPLocalConfig struct {
	RequestUrl string `mapstructure:"request_url"`
}

var NLPLocalCfg NLPLocalConfig

// InitNLPLocalConfig 初始化本地NLP模型配置
func InitNLPLocalConfig() {
	if _, err := os.Stat("./config/NLP/localConfig/localConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("./config/NLP/localConfig/localConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# http传输配置\n" +
			"# 请求发送地址（填写在你的本地模型里暴露的信息接收地址。可以为服务器地址。）\n" +
			"request_url=\"http://127.0.0.1:\""))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("localConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/NLP/localConfig") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&NLPLocalCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
