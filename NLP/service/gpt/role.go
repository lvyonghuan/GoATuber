package gpt

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// InitRole 这个函数用于获取role.cfg的角色文本信息
func InitRole() {
	file, err := os.Open("./config/NLP/GPTConfig/role.cfg")
	if err != nil {
		log.Fatalf("open config file failed: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for i := 1; ; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if i >= 2 {
			msg := strings.Split(line, ":")
			ms := &Messages{
				Role:    msg[0],
				Content: msg[1],
			}
			roleMS = append(roleMS, *ms)
		}
	}
	MS = append(MS, roleMS...)
}
