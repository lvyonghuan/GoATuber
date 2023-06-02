package pinecone

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type pineconeConfig struct {
	ApiKey      string `mapstructure:"api-key"`
	IndexName   string `mapstructure:"index"`
	ProjectName string `mapstructure:"project"`
	Environment string `mapstructure:"environment"`
}

var pineconeCfg pineconeConfig

func InitPineconeConfig() {
	if _, err := os.Stat("config/MEMORY/pinecone/pineconeConfig.cfg"); os.IsNotExist(err) {
		f, err := os.Create("config/MEMORY/pinecone/pineconeConfig.cfg")
		if err != nil {
			log.Println(err)
		}
		// 自动生成配置文件
		_, err = f.Write([]byte("# frontend.toml 配置文件\n\n" +
			"# pinecone API-KEY\n" +
			"api-key = \"xxxxx\"\n" +
			"# pinecone index 你的index名称\n" +
			"index = \"xxx\"\n" +
			"# pinecone project 你的project名称\n" +
			"project = \"xxx\"\n" +
			"# environment 你的environment名称\n" +
			"environment = \"xxx\"\n\n"))
		if err != nil {
			log.Println(err)
		}
		log.Println("配置文件不存在, 已自动生成配置文件, 请修改配置文件后再次运行程序, 5秒后退出程序...")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	viper.SetConfigName("pineconeConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/MEMORY/pinecone") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&pineconeCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
}
