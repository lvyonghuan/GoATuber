package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// ChatConfig chat模块配置信息
type ChatConfig struct {
	// 选择直播平台
	Select struct {
		bilibili bool
		YouTube  bool
	}
	Bilibili struct {
		RoomID int
	}
}

var ChatCfg ChatConfig

func InitCHATConfig() {
	if _, err := os.Stat("CHAT/chat_config.cfg"); os.IsNotExist(err) {
		f, err := os.Create("CHAT/chat_config.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# config.toml 配置文件\n\n" +
			"# 直播平台设置\n[select]\n" +
			"# B站（默认使用B站）\n" +
			"bilibili = \"1\"\n" +
			"# YouTube（暂不考虑）\n" +
			"YouTube = \"0\"\n\n" +
			"# bilibili直播间信息配置\n[bilibili]\n" +
			"#room_id = 114514\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("chat_config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./CHAT") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	err = viper.Unmarshal(&ChatCfg)
	if err != nil {
		log.Fatalf("unmarshal config failed: %v", err)
	}
}
