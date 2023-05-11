package backend

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//建立与Live2D模块的前后端交互

var done = make(chan bool, 1)
var WebsocketToNLP = make(chan bool, 1) //控制信息流入的速度

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
		if string(i) == "0" {
			WebsocketToNLP <- true
			continue
		} else {
			log.Println("错误代码：", string(i))
			WebsocketToNLP <- true
			continue
		}
	}
}
