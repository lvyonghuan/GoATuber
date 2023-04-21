package model

// Chat 弹幕消息结构体
type Chat struct {
	TimeStamp int64   //弹幕读取到的时间
	ChatName  string  //弹幕发送者名称
	message   string  //弹幕内容
	Price     float32 //sc
}

type Room struct {
	RoomID     int //房间号
	RealRoomID int //真的房间号
}
