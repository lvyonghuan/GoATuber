package proxy

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Proxy struct {
	UseProxy bool   `mapstructure:"use_proxy"`
	ProxyUrl string `mapstructure:"proxy_url"`
}

var Cfg Proxy

func InitCHATConfig() {
	if _, err := os.Stat("proxy/proxy_config.cfg"); os.IsNotExist(err) {
		f, err := os.Create("proxy/proxy_config.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# config.toml 配置文件\n\n" +
			"# 代理设置\n[proxy]\n" +
			"# openai是否走代理，默认关闭\n" +
			"use_proxy = false\n" +
			"# 代理地址\n" +
			"proxy_url = \"http://127.0.0.1:7890\"\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("proxy_config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./proxy") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("unmarshal config failed: %v", err)
	}
}
