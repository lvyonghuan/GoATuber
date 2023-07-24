package get_live2d_model_info

import (
	"log"
	"os"
	"regexp"
)

type modelConfig struct {
	Mouth float64
}

func GetModelName() string {
	files, err := os.ReadDir("dist/live2d")
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
