package frontend

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	config := r.Group("/config")
	{
		config.GET("/configInfo")
		config.PUT("/NLP")
		config.PUT("/CHAT")
		config.PUT("/proxy")
	}
	r.Run()
}
