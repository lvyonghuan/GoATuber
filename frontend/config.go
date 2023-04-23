// Package frontend 用于读写config文件
package frontend

import (
	config2 "GoTuber/CHAT/config"
	"GoTuber/NLP/config"
	proxy2 "GoTuber/proxy"

	"github.com/gin-gonic/gin"
)

const (
	NLPConfig   string = "NLP/config.cfg"
	CHATConfig  string = "CHAT/config.cfg"
	ProxyConfig string = "proxy/config.cfg"
)

func GetConfigInfo(c gin.Context) {
	nlp := config.GptConfig{}
	chat := config2.ChatConfig{}
	proxy := proxy2.Proxy{}
}
