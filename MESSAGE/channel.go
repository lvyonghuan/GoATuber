package MESSAGE

import (
	"GoTuber/MESSAGE/model"
	"GoTuber/NLP/service"
)

var ChatToFilter = make(chan model.Chat, 2) //要是出问题了就改大，但是应该没问题了
var SelectToNLP = make(chan model.Chat, 1)

func GetMessage() {
	for {
		select {
		case msg := <-ChatToFilter:
			go FILTER(msg)
		case msg := <-SelectToNLP:
			service.GetMessageFromChat(msg)
		default:
			continue
		}
	}
}
