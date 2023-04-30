// Package sensitive 因为循环调用的问题，把ai的filter单独放在了这里
package sensitive

import (
	"log"
	"math/rand"
	"time"
)

//func AIFilter(msg string) {
//	filter := New()
//	err := filter.LoadWordDict("MESSAGE/filter/dict/dict.txt")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	isValid, _ := filter.Trie.Validate(msg)
//	if !isValid {
//		log.Println("filter!")
//		msg = "filter!"
//		return
//	}
//	service.GetMessage(msg)
//}

type OutPut struct {
	Msg string
	Id  int
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
	Msg.Id = randomFiveDigits()
}

// 随机数生成器，给传出消息打标签，用于信息的同步
func randomFiveDigits() int {
	rand.Seed(time.Now().UnixNano())  // 设置随机数种子为当前时间戳
	min, max := 10000, 99999          // 定义随机数范围
	return rand.Intn(max-min+1) + min // 生成并返回随机数
}
