// Package backend 前后端交互入口：借助websocket进行
package backend

import (
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MOOD"
	"math/rand"
	"time"
)

var OutPutChan = make(chan OutMessage, 1)

type OutMessage struct {
	Voice      string `json:"voice"`
	VType      int    `json:"VType"`      //voice格式type，1表示http，2表示base64编码,3为二进制（编码成base64）
	Act        string `json:"act"`        //数组名称
	Movement   string `json:"movement"`   //动作，全身的
	Expression string `json:"expression"` //表情，脸部的
}

func GetMessage(msg *sensitive.OutPut) {
	var putOut OutMessage
	putOut.Voice = msg.Voice
	putOut.VType = msg.VType
	rand.Seed(time.Now().UnixNano()) //如果相同情绪有多个动作......摆了，随机选一个
	if msg.Mood == "happy" {
		getAction(&putOut, msg.Mood)
	} else if msg.Mood == "mad" {
		getAction(&putOut, msg.Mood)
	} else if msg.Mood == "sad" {
		getAction(&putOut, msg.Mood)
	} else if msg.Mood == "disgust" {
		getAction(&putOut, msg.Mood)
	} else if msg.Mood == "surprise" {
		getAction(&putOut, msg.Mood)
	} else if msg.Mood == "fear" {
		getAction(&putOut, msg.Mood)
	} else if msg.Mood == "health" {
		getAction(&putOut, msg.Mood)
	}
	OutPutChan <- putOut
}

// 动作生成
func getAction(putOut *OutMessage, mood string) {
	action, ok := MOOD.MoodAct[mood+"1"] //判断动作是否存在，如果存在即可取出动作

	if ok {
		index := randomIndex(action)
		putOut.Act = action[index].Event
		putOut.Movement = action[index].Act
	}
	action, ok = MOOD.MoodAct[mood+"2"]
	if ok {
		index := randomIndex(action)
		putOut.Act = action[index].Event
		putOut.Expression = action[index].Act
	}
}

// 随机生成一个索引
func randomIndex(movement []MOOD.Mood) int {
	l := len(movement)
	randomIndex := rand.Intn(l)
	return randomIndex
}
