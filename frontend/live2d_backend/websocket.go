package backend

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//建立与Live2D模块的前后端交互

//TODO:有时间迟早得重构整个项目。1.把各个http接口调用程序重新归类;2.改成context控制。不得不说这两个月还是学了点东西。

var done = make(chan bool, 1)
var WebsocketToNLP = make(chan bool, 1)    //控制信息流入的速度
var WebsocketToSpeech = make(chan bool, 1) // 控制语音对话交流

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Start(c *gin.Context) {
	conn, err := Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("与web建立websocket连接失败：", err)
		return
	}
	WebsocketToNLP <- true //初始化NLP信息读取模块
	go write(conn)
	go read(conn)
}

func write(conn *websocket.Conn) {
	defer func() {
		log.Println("连接断开")
	}()
	for {
		select {
		case msg := <-OutPutChan:
			// 将OutMessage结构体转换为JSON格式
			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				log.Println(err)
				continue
			}
			// 发送JSON消息到WebSocket连接
			err = conn.WriteMessage(websocket.TextMessage, jsonMsg)
			if err != nil {
				log.Println(err)
				continue
			}
		case <-done:
			return
		}
	}
}

func read(conn *websocket.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
	for {
		_, i, err := conn.ReadMessage()
		if err != nil {
			done <- true
			break
		}
		if string(i) == "0" { //朗读结束标志
			WebsocketToNLP <- true
			continue
		} else if string(i) == "1" { //收到语音消息
			WebsocketToSpeech <- true
			continue
		} else {
			log.Println("错误代码：", string(i))
			WebsocketToNLP <- true
			continue
		}
	}
}
