package MESSAGE

//处理直播消息获取模块与对话生成模块之间的消息传递逻辑

import (
	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/service"
)

var ChatToFilter = make(chan model.Chat, 50) //在洁宝直播间做了测试，评价是不如改大
var FilterToNLP = make(chan model.Chat, 1)

func GetMessage() {
	for {
		select {
		//读取到的消息流向过滤器进行过滤
		case msg := <-ChatToFilter:
			if messageCfg.IsGetChat { //如果不启用弹幕消息，则直接放掉
				go FILTER(msg)
			}
		//消息从过滤器流向NLP模块
		case msg := <-FilterToNLP:
			service.GetMessageFromChat(msg)
		}
	}
}
