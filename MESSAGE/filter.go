package MESSAGE

import (
	"GoTuber/MESSAGE/filter"
	"GoTuber/MESSAGE/model"
	"log"
)

// FILTER 过滤器，基于importcjj/sensitive实现，证书在filter目录下。使用协程进行过滤后，汇总到一条队列里。
func FILTER(msg model.Chat) {
	filter := sensitive.New()
	err := filter.LoadWordDict("MESSAGE/filter/dict/dict.txt", 0)
	if err != nil {
		log.Println(err)
		return
	}
	isValid, _ := filter.Trie.Validate(msg.Message)
	if !isValid {
		//TODO：一些操作，比如塞进某个数据库。
		return
	}
	FilterToNLP <- msg
}
