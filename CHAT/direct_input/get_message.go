package direct_input

import (
	"GoTuber/MESSAGE"
	"GoTuber/MESSAGE/model"
	"bufio"
	"log"
	"os"
	"strings"
)

func GetMessage() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ms := strings.TrimSpace(scanner.Text())
			if ms == "" {
				continue // 忽略空行
			}
			var msg model.Chat
			msg.Message = ms
			msg.ChatName = "管理员"
			MESSAGE.ChatToFilter <- msg
		}
		if err := scanner.Err(); err != nil {
			log.Println("输入失败,", err)
		}
	}
}
