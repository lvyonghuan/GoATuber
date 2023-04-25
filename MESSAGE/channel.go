package MESSAGE

import (
	sensitive "GoTuber/MESSAGE/filter_and_select/filter"
	"GoTuber/MESSAGE/model"
)

var ChatToFilter = make(chan model.Chat, 2) //要是出问题了就改大，但是应该没问题了

func GetMessage() {
	for {
		select {
		case msg := <-ChatToFilter:
			go sensitive.FILTER(msg)
		}
	}
}
