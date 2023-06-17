package service

import (
	"GoTuber/frontend/live2d_backend"
	"container/list"
	"log"
	"strconv"
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

var MsgMu sync.Mutex //这个锁在接收到消息的时候解开，防止程序在没有弹幕的时候进入死循环的不断查询状态

//使用三个管道进行循环通信，控制流程，减少性能消耗。

var GetToChooseFlag = make(chan bool, 1)
var ChooseToReadFlag = make(chan bool, 1)
var ReadToGetFlag = make(chan bool, 1)

func GetMessageFromChat(msg model.Chat) {
	if !MsgMu.TryLock() {
		MsgMu.Unlock()
	}
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
				MsgMu.Lock()
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
					HandelMsg.Uid = strconv.Itoa(scMsgList.Front().Value.(*model.Chat).Uid)
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
					HandelMsg.Uid = strconv.Itoa(normalMsgList.Front().Value.(*model.Chat).Uid)
					HandelMsg.IsUse = true
					HandelMsg.Mu.Unlock()
					norMu.Lock()
					normalMsgList.Remove(normalMsgList.Front())
					norMu.Unlock()
					ChooseToReadFlag <- true
					continue
				} else {
					GetToChooseFlag <- true //没弹幕的时候，防止管道堵塞（msgMu应该把这个状况给避免了...希望如此，不然这里会发生死循环从而占用cpu）
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
			log.Println("正在生成文本......")
			if config.NLPCfg.Nlp.UseGPT {
				gpt.GenerateTextByOpenAI(&HandelMsg)
			} else if config.NLPCfg.Nlp.UseAzureGPT {
				gpt.GenerateTextByAzureOpenAI(&HandelMsg)
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
