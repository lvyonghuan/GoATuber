package sensitive

import (
	"bytes"
	"io"
	"net/http"
)

//这是一个预留的接口。使用http协议与其他敏感词过滤方案通信。当时写filter的时候只想到字典过滤了，算我的过失吧。
//输入字符串，作为过滤信息。返回布尔量，用于判断是否通过了过滤。

func UserOtherFilter(message string) bool {
	req, _ := http.NewRequest("GET", FilterCfg.RequestUrl, bytes.NewBuffer([]byte(message)))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp == nil {
		return false
	}
	body, _ := io.ReadAll(resp.Body)
	state := string(body)
	if state == "1" {
		return true
	} else {
		return false
	}
}
