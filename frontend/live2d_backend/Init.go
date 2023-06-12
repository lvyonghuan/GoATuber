package backend

import (
	"GoTuber/frontend/live2d_backend/get_live2d_model_info"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func Init() {
	r := gin.Default()
	r.GET("/live2d", Start)
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
	r.GET("/get", getModelInfo)
	r.Run(":9000") //服务在本地9000端口运行
}

//ModelInfo初始化内容↓

type modelInfo struct {
	Name  string  `json:"name"`
	Mouth float64 `json:"mouth"`
}

func getModelInfo(c *gin.Context) {
	name := get_live2d_model_info.GetModelName()
	_ = get_live2d_model_info.GetModelConfig()
	if name == "" {
		return
	}
	c.JSON(http.StatusOK, name)
}
