package gpt

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var roleLine int

// InitRole 这个函数用于获取role.cfg的角色文本信息
func InitRole() {
	file, err := os.Open("./config/NLP/GPTConfig/role.cfg")
	if err != nil {
		log.Fatalf("open config file failed: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for roleLine = 1; ; roleLine++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if roleLine >= 2 {
			msg := strings.Split(line, ":")
			ms := &RequestMessages{
				Role:    msg[0],
				Content: msg[1],
				Name:    "system",
			}
			roleMS = append(roleMS, *ms)
		}
	}
	roleLine -= 1
	MS = append(MS, roleMS...)
}
