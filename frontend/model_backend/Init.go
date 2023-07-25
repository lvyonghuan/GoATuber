package backend

import (
	"net/http"
	"os"
	"strings"

	"GoTuber/frontend/model_backend/get_model_info"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.GET("/connect", Start) //启动websocket连接
	r.Use(static.Serve("/", static.LocalFile("dist", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("./dist/index.html")
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
	r.GET("/get", getModelInfo) //获取模型类型信息
	r.Run(":9000")              //服务在本地9000端口运行
}

//ModelInfo初始化内容↓

func getModelInfo(c *gin.Context) {
	const (
		nil    = 0
		live2d = 1 //这两个放这里是提醒我的
		vrm    = 2
	)
	var info get_model_info.ModelInfo
	info.GetModelName()
	//大概不会有这种情况
	if info.Type == nil {
		return
	}
	c.JSON(http.StatusOK, info)
}
