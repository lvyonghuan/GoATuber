// Package sensitive 因为循环调用的问题，把ai的filter单独放在了这里
package sensitive

import (
	"log"
	"sync"
)

type OutPut struct {
	Msg   string
	Mu    sync.Mutex
	Voice string
	Mood  string
	VType int //约定：1，http格式;2，base64格式
}

func (OutPut) AIFilter(Msg *OutPut) {
	msg := Msg.Msg
	filter := New()
	err := filter.LoadWordDict("MESSAGE/filter/dict/dict.txt", 0)
	if err != nil {
		log.Println(err)
		return
	}
	isValid, _ := filter.Trie.Validate(msg)
	if !isValid {
		log.Println("filter!")
		Msg.Msg = "filter!"
		return
	}
}
