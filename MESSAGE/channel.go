package MESSAGE

import (
	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/service"
)

var ChatToFilter = make(chan model.Chat, 50) //在洁宝直播间做了测试，评价是不如改大
var FilterToNLP = make(chan model.Chat, 1)

func GetMessage() {
	for {
		select {
		case msg := <-ChatToFilter:
			go FILTER(msg)
		case msg := <-FilterToNLP:
			service.GetMessageFromChat(msg)
		default:
			continue
		}
	}
}
