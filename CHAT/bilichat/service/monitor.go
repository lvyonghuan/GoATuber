package service

import (
	config2 "GoTuber/CHAT/config"
	"log"
	"sync"
	"time"
)

const (
	danMuMsgBufCap = 256
)

type Monitor struct {
	servers []*ChatServer
	group   sync.WaitGroup
}

func NewMonitor(c config2.ChatConfig) *Monitor {
	m := &Monitor{
		servers: make([]*ChatServer, 0),
	}
	chatServer, err := GetChatServer(c.Bilibili.RoomID)
	if err != nil {
		log.Println("获取弹幕服务器失败：", err)
		return nil
	}
	m.servers = append(m.servers, chatServer)
	return m
}

func work(chat *ChatServer, group *sync.WaitGroup) {
	out := make(chan Message, ChanBufSize*2) //两倍缓冲
	go chat.ReceiveMsg(out)
	for {
		msg, ok := <-out
		if !ok {
			group.Done()
			return
		}
		switch m := msg.(type) {
		case *DanMuMessage:
			log.Println(m)
			//case *SuperChatMessage:
			//
			//case *GiftMessage:
			//	ifInsertError(d.insertGiftMsg(*r, m))
			//case *GuardMessage:
			//	ifInsertError(d.insertGuardMsg(*r, m))
			//case *EntryMessage:
			//	ifInsertError(d.insertEntryMsg(*r, m))
			//case *RoomFansMessage:
			//	ifInsertError(d.insertFansMsg(*r, m))
		}
	}
}

func (m *Monitor) Start() {
	for _, c := range m.servers {
		err := c.Connect()
		if err != nil {
			log.Println("连接直播间失败:", err)
			return
		}
		m.group.Add(1)
		go work(c, &(m.group))
		time.Sleep(time.Second)
	}
}

func (m *Monitor) Stop() {
	log.Println("程序退出...")
	for _, c := range m.servers {
		c.Disconnect()
	}
	m.group.Wait()
}
