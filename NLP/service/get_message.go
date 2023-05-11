package service

import (
	"GoTuber/frontend/live2d_backend"
	"container/list"
	"sync"
	"time"

	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/config"
	"GoTuber/NLP/service/gpt"
)

var HandelMsg model.Msg
var scMsgList list.List     //氪金消息队列
var normalMsgList list.List //普通消息队列，定长，队首自动弹出
var scMu sync.Mutex         //氪金消息锁
var norMu sync.Mutex        //普通消息锁

//使用三个管道进行循环通信，控制流程，减少性能消耗。

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
	go func() {
		for {
			select {
			case <-ReadToGetFlag:
				GetToChooseFlag <- true
			}
		}
	}()
}

// ChooseMessage 消息选择器
func ChooseMessage() {
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
				} else {
					GetToChooseFlag <- true //没弹幕的时候，防止管道堵塞
				}
			}
		}
	}
}

// HandelMessage 将消息传送给具体的处理模块
func HandelMessage() {
	for {
		<-ChooseToReadFlag
		<-backend.WebsocketToNLP
		if HandelMsg.IsUse {
			if config.NLPCfg.Nlp.UseGPT {
				gpt.GenerateText(&HandelMsg)
			} else if config.NLPCfg.Nlp.UseOther {
				//TODO：以后再说
			}
			HandelMsg.Mu.Lock()
			HandelMsg.IsUse = false
			HandelMsg.Mu.Unlock()
			ReadToGetFlag <- true
		}
	}
}
