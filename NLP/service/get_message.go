package service

import (
	"GoTuber/MESSAGE/model"
	"container/list"
	"log"
	"sync"
	"time"
)

type Msg struct {
	Msg   string
	Name  string
	IsUse bool
	Mu    sync.Mutex
}

var HandelMsg Msg
var scMsgList list.List     //氪金消息队列
var normalMsgList list.List //普通消息队列，定长，队首自动弹出
var scMu sync.Mutex         //氪金消息锁
var norMu sync.Mutex        //普通消息锁

func GetMessageFromChat(msg model.Chat) {
	//TODO:完善sc队列机制，用户规定sc阈值
	if msg.Price > 0 {
		scMu.Lock()
		scMsgList.PushBack(&msg)
		scMu.Unlock()
	} else {
		norMu.Lock()
		normalMsgList.PushBack(&msg)
		norMu.Unlock()
	}
}

// ChooseMessage 消息选择器
func ChooseMessage() {
	log.Println("start!")
	go func() {
		for {
			if normalMsgList.Len() > 50 {
				norMu.Lock()
				normalMsgList.Init()
				norMu.Unlock()
			}
		}
	}()
	for {
		if !HandelMsg.IsUse {
			//优先选择sc队列中消息
			if scMsgList.Len() > 0 {
				scMu.Lock()
				HandelMsg.Mu.Lock()
				HandelMsg.Msg = scMsgList.Front().Value.(*model.Chat).Message
				HandelMsg.Name = scMsgList.Front().Value.(*model.Chat).ChatName
				HandelMsg.IsUse = true
				scMsgList.Remove(scMsgList.Front())
				HandelMsg.Mu.Unlock()
				scMu.Unlock()
				continue
			} else if normalMsgList.Len() > 0 {
				norMu.Lock()
				HandelMsg.Mu.Lock()
				HandelMsg.Msg = normalMsgList.Back().Value.(*model.Chat).Message
				HandelMsg.Name = normalMsgList.Back().Value.(*model.Chat).ChatName
				HandelMsg.IsUse = true
				normalMsgList.Remove(normalMsgList.Back())
				HandelMsg.Mu.Unlock()
				norMu.Unlock()
				continue
			}
		}
	}
}

// Read test
func Read() {
	for {
		if HandelMsg.IsUse {
			HandelMsg.Mu.Lock()
			time.Sleep(time.Second * 5)
			log.Println(normalMsgList.Len())
			log.Println(HandelMsg.Name, HandelMsg.Msg)
			HandelMsg.IsUse = false
			HandelMsg.Mu.Unlock()
		}
	}
}
