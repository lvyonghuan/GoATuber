package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type NLPConfig struct {
	Nlp struct {
		UseGPT      bool `mapstructure:"use_gpt"`
		UseAzureGPT bool `mapstructure:"use_azure_gpt"`
		UseOther    bool `mapstructure:"use_other"`
	}
}

var NLPCfg NLPConfig

// InitNLPConfig 初始化NLP模块配置
func InitNLPConfig() {
	if _, err := os.Stat("config/NLP/NLPConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/NLP/NLPConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# NLP模块配置\n[NLPConfig]\n" +
			"# 是否使用GPT模型\n" +
			"use_gpt = true\n" +
			"# 是否使用azure GPT\n" +
			"use_azure_gpt = false\n" +
			"# 是否使用其他模型\n" +
			"use_other = false\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("NLPConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/NLP") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&NLPCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
