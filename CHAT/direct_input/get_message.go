package direct_input

import (
	"GoTuber/MESSAGE"
	"GoTuber/MESSAGE/model"
	"fmt"
	"log"
)

func GetMessage() {
	go func() {
		for {
			var ms string
			_, err := fmt.Scanln(&ms)
			if err != nil {
				log.Println("输入失败，", err)
				return
			}
			var msg model.Chat
			msg.Message = ms
			msg.ChatName = "访客"
			MESSAGE.ChatToFilter <- msg
		}
	}()
}
