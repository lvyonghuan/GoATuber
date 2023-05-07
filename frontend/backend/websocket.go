package backend

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var done = make(chan bool, 1)
var WebsocketToNLP = make(chan bool, 1) //控制信息流入的速度

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
	r.Use(static.Serve("/", static.LocalFile("dist", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("dist/index.html")
			if err != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})
	r.Run(":9000") //服务在本地9000端口运行
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
			continue
		}
	}
}
