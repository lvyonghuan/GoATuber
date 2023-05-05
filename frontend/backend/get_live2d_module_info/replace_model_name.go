package get_live2d_module_info

import (
	"log"
	"os"
	"regexp"
)

func ReplaceName() {
	jsBytes, err := os.ReadFile("dist/js/app.862a8329.js")
	if err != nil {
		log.Fatal("读取js文件失败，错误信息：", err)
		return
	}
	jsStr := string(jsBytes)

	// 正则表达式匹配字段的前后文
	re := regexp.MustCompile(`this\.model4=await a\._Y\.from\("\.\/model\/(.+?)\.model3\.json"`)
	matches := re.FindStringSubmatch(jsStr)
	if len(matches) < 2 {
		log.Fatal("无法匹配到字段")
		return
	}
	// 将字段替换为你想要的字符串
	newJsStr := re.ReplaceAllString(jsStr, "this.model4=await a._Y.from(\"./model/"+Live2dCfg.Live2d.Name+".model3.json\"")
	err = os.WriteFile("dist/js/app.862a8329.js", []byte(newJsStr), 0644)
	if err != nil {
		log.Fatal("更新js失败，错误信息：", err)
	}
}
