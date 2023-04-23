package service

// 数据类型
const (
	verPlain  = 0 //普通文本，utf-8编码
	verInt    = 1 //控制通信消息
	verZlib   = 2 // zlib格式压缩，已弃用
	verBrotli = 3 //brotli压缩
)

// 操作码
const (
	opHeartbeat      = 2 //发送心跳包
	opHeartbeatReply = 3 //服务端回应心跳包
	opMessage        = 5 //弹幕消息等
	opEnterRoom      = 7 //进入直播间
	opEnterRoomReply = 8 //进入直播间成功
)

// Cmd
const (
	CmdEntryEffect      = "ENTRY_EFFECT"       //舰长进场消息
	CmdSendGift         = "SEND_GIFT"          //投喂礼物
	CmdComboSend        = "COMBO_SEND"         //礼物连击
	CmdDanMuMSG         = "DANMU_MSG"          //弹幕
	CmdUserToastMsg     = "USER_TOAST_MSG"     //续费舰长
	CmdSuperChatMessage = "SUPER_CHAT_MESSAGE" //sc
)
