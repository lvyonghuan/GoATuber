package sensitive

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type filterConfig struct {
	UseDict    bool   `mapstructure:"use_dict"`
	UseOther   bool   `mapstructure:"use_other_filter"`
	RequestUrl string `mapstructure:"request_url"`
}

var FilterCfg filterConfig

func InitConfig() {
	if _, err := os.Stat("config/MESSAGE/filter/filter.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/MESSAGE/filter/filter.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# 使用本地字典（误伤率高）\n" +
			"use_dict = true\n" +
			"# 其他过滤方式\n" +
			"use_other_filter = false\n" +
			"# 发送请求地址（只启用本地字典不用管）\n" +
			"request_url = \"\"\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("filter.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/MESSAGE/filter") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&FilterCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
