package backend

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var done = make(chan bool, 1)

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Init() {
	r := gin.Default()
	r.GET("/live2d", Start)
	r.Run(":9000") //服务在本地9000端口运行
}

func Start(c *gin.Context) {
	conn, err := Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("与web建立websocket连接失败：", err)
		return
	}
	go write(conn)
	go watch(conn)
}

func write(conn *websocket.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
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
}

func watch(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			done <- true
			return
		}
	}
}
