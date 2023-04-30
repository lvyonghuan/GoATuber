package model

import "sync"

// Chat 弹幕（聊天）消息结构体
type Chat struct {
	TimeStamp int64   //弹幕读取到的时间
	ChatName  string  //弹幕发送者名称
	Message   string  //弹幕内容
	Price     float32 //sc
	Uid       int     //B站用户uid，不知道油管有没有
}

type Msg struct {
	Msg   string
	Name  string
	IsUse bool
	Mu    sync.Mutex
}
