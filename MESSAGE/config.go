package MESSAGE

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type messageConfig struct {
	IsGetChat bool `mapstructure:"is_get_chat"`
}

var messageCfg messageConfig

func InitConfig() {
	if _, err := os.Stat("config/MESSAGE/message.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/MESSAGE/message.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# 是否读取弹幕消息（专门的对话直播可以改成false，避免打扰）\n" +
			"is_get_chat = true\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("message.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/MESSAGE/") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&messageCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
