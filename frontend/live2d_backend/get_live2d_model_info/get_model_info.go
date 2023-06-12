package get_live2d_model_info

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"regexp"
)

type modelConfig struct {
	Mouth float64
}

func GetModelName() string {
	files, err := os.ReadDir("dist/model")
	if err != nil {
		log.Fatalf("读取文件错误！错误信息：%v", err)
	}
	for _, file := range files {
		fileName := file.Name()
		match, err := regexp.Match("^.*\\.model3\\.json$", []byte(fileName))
		if err != nil {
			log.Println("live2d初始化警告：正则表达式匹配出错！", err)
		}
		if match {
			return fileName
		}
	}
	log.Println("未找到模型文件！特别说明：该项目目前只支持model3.json后缀的模型文件")
	return ""
}

func GetModelConfig() float64 {
	var modelCfg modelConfig
	viper.SetConfigName("ModelConfig.cfg")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./frontend/live2d_backend/get_live2d_model_info") // 指定查找配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read frontend failed: %v", err)
	}
	err = viper.Unmarshal(&modelCfg)
	if err != nil {
		log.Fatalf("unmarshal frontend failed: %v", err)
	}
	return modelCfg.Mouth
}
