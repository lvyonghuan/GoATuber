// Package websocket 前后端交互：借助websocket进行
package websocket

import (
	sensitive "GoTuber/MESSAGE/filter"
	"log"
)

func GetMessage(msg *sensitive.OutPut) {
	log.Println(msg)
}
