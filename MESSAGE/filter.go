package MESSAGE

import (
	"log"

	"GoTuber/MESSAGE/filter"
	"GoTuber/MESSAGE/model"
)

// FILTER 过滤器
func FILTER(msg model.Chat) {
	isValid := handelMessage(msg.Message)
	if !isValid {
		return
	}
	FilterToNLP <- msg
}

func handelMessage(message string) bool {
	if sensitive.FilterCfg.UseDict == true {
		filter := sensitive.New()
		err := filter.LoadWordDict("config/MESSAGE/filter/dict/dict.txt", 0)
		if err != nil {
			log.Println(err)
			return false
		}
		isValid, _ := filter.Trie.Validate(message)
		return isValid
	} else if sensitive.FilterCfg.UseOther == true {
		return sensitive.UserOtherFilter(message)
	}
	return true //不用filter
}
