package gpt

import "log"

//清理消息记录，减少token使用量

// 清除每次附带在消息上的记忆
func cleanMemoryMessage() {
	MS = append(MS[:roleLine], MS[roleLine+1:]...)
}

// 清除所有的消息
func cleanAllMessage() {
	MS = MS[:roleLine]
}

// 队列式的清除消息。达到token限制之后，清除队首的消息即可。
func cleanFirstMessage() {
	MS = append(MS[:roleLine], MS[roleLine+1:]...)
	log.Println(MS[:roleLine-1], MS[roleLine+1:])
}
