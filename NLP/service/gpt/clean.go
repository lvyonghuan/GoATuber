package gpt

//清理消息记录，减少token使用量

// 清除每次附带在消息上的记忆
func cleanMemoryMessage() {
	MS = append(MS[:roleLine], MS[roleLine+1:]...)
}

// 清除所有的消息
func cleanAllMessage() {
	MS = MS[:roleLine]
}
