package MEMORY

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type memoryConfig struct {
	IsUse bool `mapstructure:"is_use"` //是否启用记忆模块
}

var MemoryCfg memoryConfig

func InitMemoryConfig() {
	if _, err := os.Stat("config/MEMORY/MemoryConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/MEMORY/MemoryConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# 是否启用记忆模块（启用可能导致大量token消耗）\n" +
			"is_use = false\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("MemoryConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/MEMORY") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&MemoryCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
