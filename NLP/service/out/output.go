package out

import (
	"GoTuber/MESSAGE/filter"
	"GoTuber/MOOD"
	"GoTuber/SPEECH/service"
	"GoTuber/frontend/live2d_backend"
)

// PutOutMsg 统一将NLP模块的消息向外传递
func PutOutMsg(msg *sensitive.OutPut) {
	msg.AIFilter(msg)
	MOOD.GetMessage(msg)
	service.GetMessage(msg)
	backend.GetMessage(msg)
}
