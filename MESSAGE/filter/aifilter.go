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
		msg = "filter!"
		return
	}
}
