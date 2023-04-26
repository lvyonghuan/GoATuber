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

var GetToChooseFlag = make(chan bool, 1)
var ChooseToReadFlag = make(chan bool, 1)
var ReadToGetFlag = make(chan bool, 1)

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
	select {
	case <-ReadToGetFlag:
		GetToChooseFlag <- true
	}
}

// ChooseMessage 消息选择器
func ChooseMessage() {
	//TODO:用户手动调整sleep时间
	log.Println("start!")
	go func() {
		for {
			if normalMsgList.Len() > 5 {
				norMu.Lock()
				normalMsgList.Init()
				norMu.Unlock()
			}
			time.Sleep(60 * time.Second)
		}
	}()
	for {
		select {
		case <-GetToChooseFlag:
			if !HandelMsg.IsUse {
				//优先选择sc队列中消息
				if scMsgList.Len() > 0 {
					HandelMsg.Mu.Lock()
					HandelMsg.Msg = scMsgList.Front().Value.(*model.Chat).Message
					HandelMsg.Name = scMsgList.Front().Value.(*model.Chat).ChatName
					HandelMsg.IsUse = true
					HandelMsg.Mu.Unlock()
					scMu.Lock()
					scMsgList.Remove(scMsgList.Front())
					scMu.Unlock()
					ChooseToReadFlag <- true
					continue
				} else if normalMsgList.Len() > 0 {
					HandelMsg.Mu.Lock()
					HandelMsg.Msg = normalMsgList.Front().Value.(*model.Chat).Message
					HandelMsg.Name = normalMsgList.Front().Value.(*model.Chat).ChatName
					HandelMsg.IsUse = true
					HandelMsg.Mu.Unlock()
					norMu.Lock()
					normalMsgList.Remove(normalMsgList.Front())
					norMu.Unlock()
					ChooseToReadFlag <- true
					continue
				}
			}
		}
	}
}

// Read test
func Read() {
	for {
		select {
		case <-ChooseToReadFlag:
			if HandelMsg.IsUse {
				time.Sleep(time.Second * 5)
				log.Println(HandelMsg.Name, HandelMsg.Msg)
				HandelMsg.Mu.Lock()
				HandelMsg.IsUse = false
				HandelMsg.Mu.Unlock()
				ReadToGetFlag <- true
			}
		}
	}
}
