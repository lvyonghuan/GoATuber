package get_model_info

import (
	"log"
	"os"
	"regexp"
)

type ModelInfo struct {
	Type int    `json:"type"`
	Name string `json:"name"`
}

func (info *ModelInfo) GetModelName() {
	const (
		live2d = 1
		vrm    = 2
	)
	files, err := os.ReadDir("dist/vrm")
	if err != nil {
		log.Fatalf("读取文件错误！错误信息：%v", err)
	}
	for _, file := range files {
		fileName := file.Name()
		//进行vrm模型匹配
		matchVRM, err := regexp.Match(".*\\.vrm$", []byte(fileName))
		if err != nil {
			log.Println("模型初始化警告：正则表达式匹配出错！", err)
		}
		if matchVRM {
			info.Type = vrm
			info.Name = fileName
			return
		}

		//进行live2d匹配
		matchLive2d, err := regexp.Match("^.*\\.model3\\.json$", []byte(fileName))
		if err != nil {
			log.Println("模型初始化警告：正则表达式匹配出错！", err)
		}
		if matchLive2d {
			info.Type = live2d
			info.Name = fileName
			return
		}
	}
	log.Fatalf("未找到模型文件！")
	return
}
