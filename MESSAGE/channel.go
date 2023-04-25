package MESSAGE

import (
	"GoTuber/MESSAGE/model"
	"log"
)

var ChatToFilter = make(chan model.Chat, 2) //要是出问题了就改大，但是应该没问题了

func GetMessage() {
	for {
		select {
		case msg := <-ChatToFilter:
			log.Println(msg)
		}
	}
}
