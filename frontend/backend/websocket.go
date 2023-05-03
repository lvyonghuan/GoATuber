package backend

import (
	"encoding/json"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
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
	//r.LoadHTMLFiles("dist/index.html", "dist/index.html")
	r.GET("/live2d", Start)
	r.Use(static.Serve("/start", static.LocalFile("dist", true)))
	//r.Use(static.Serve("/dist", static.LocalFile("dist", true))) // 添加此行代码
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
			c.Writer.Write((content))
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
	//go read(conn)
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

//func read(conn *websocket.Conn) {
//	defer func() {
//		err := conn.Close()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//	}()
//	for {
//		log.Println("y")
//		_, i, err := conn.ReadMessage()
//		log.Println(i)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		if string(i) == "0" {
//			log.Println("yes")
//			//TODO：加入反馈
//			continue
//		} else {
//			log.Println("错误代码：", string(i))
//			continue
//		}
//	}
//}

func watch(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			done <- true
			return
		}
	}
}
