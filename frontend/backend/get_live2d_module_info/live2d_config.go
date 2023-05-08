package get_live2d_module_info

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type live2d struct {
	Live2d struct {
		Name string `mapstructure:"name"`
	}
}

var Live2dCfg live2d

func InitLive2DConfig() {
	if _, err := os.Stat("frontend/backend/get_live2d_module_info/live2d_config.cfg"); os.IsNotExist(err) {
		f, err := os.Create("frontend/backend/get_live2d_module_info/live2d_config.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# live2d配置设置\n[live2d]\n" +
			"# 模型名称（model3.json）\n" +
			"name = \"\"\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("live2d_config.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./frontend/backend/get_live2d_module_info") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&Live2dCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
	ReplaceName()
}
