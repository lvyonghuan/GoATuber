package MESSAGE

import (
	sensitive "GoTuber/MESSAGE/filter_and_select/filter"
	"GoTuber/MESSAGE/model"
	"log"
)

// FILTER 过滤器，基于importcjj/sensitive实现，证书在filter目录下。使用协程进行过滤后，汇总到一条队列里。
func FILTER(msg model.Chat) {
	filter := sensitive.New()
	err := filter.LoadWordDict("MESSAGE/filter_and_select/filter/dict/dict.txt")
	if err != nil {
		log.Println(err)
		return
	}
	isValid, _ := filter.Trie.Validate(msg.Message)
	if !isValid {
		//TODO：一些操作，比如塞进某个数据库。
		return
	}
	SelectToNLP <- msg
}
