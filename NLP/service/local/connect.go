package local

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/config"
	"GoTuber/NLP/service/out"
	backend "GoTuber/frontend/model_backend"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Connect 该模块会按照config文件配置的端口号发送request。
func Connect(msg *model.Msg) {
	requestMessage := request{
		Message:  msg.Msg,
		Username: msg.Name,
	}
	requestJson, err := json.Marshal(requestMessage)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println("json序列化错误：", err)
		return
	}
	req, _ := http.NewRequest("GET", config.NLPLocalCfg.RequestUrl, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		backend.WebsocketToNLP <- true
		log.Println("请求错误：", err)
		return
	}
	defer resp.Body.Close()
	if resp == nil {
		backend.WebsocketToNLP <- true
		log.Println("response is nil")
		return
	}
	body, _ := io.ReadAll(resp.Body)
	var response response
	err = json.Unmarshal(body, &response)
	if response.Type == 0 {
		log.Println("生成错误，错误消息：", response.Message)
		backend.WebsocketToNLP <- true
		return
	}
	var ms sensitive.OutPut
	ms.Msg = response.Message
	out.PutOutMsg(&ms)
}
