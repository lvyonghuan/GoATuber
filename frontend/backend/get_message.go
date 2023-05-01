// Package backend 前后端交互入口：借助websocket进行
package backend

import (
	sensitive "GoTuber/MESSAGE/filter"
)

var OutPutChan = make(chan OutMessage, 1)

type OutMessage struct {
	Voice string `json:"voice"`
	Mood  int    `json:"mood"`
}

//约定：happy,0 mad,1 sad,2 disgust,3 surprise,4 fear,5 health,6

func GetMessage(msg *sensitive.OutPut) {
	var putOut OutMessage
	putOut.Voice = msg.Voice
	if msg.Mood == "happy" {
		putOut.Mood = 0
	} else if msg.Mood == "mad" {
		putOut.Mood = 1
	} else if msg.Mood == "sad" {
		putOut.Mood = 2
	} else if msg.Mood == "disgust" {
		putOut.Mood = 3
	} else if msg.Mood == "surprise" {
		putOut.Mood = 4
	} else if msg.Mood == "fear" {
		putOut.Mood = 5
	} else if msg.Mood == "health" {
		putOut.Mood = 6
	}
	OutPutChan <- putOut
}
