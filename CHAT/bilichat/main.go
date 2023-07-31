package bilichat

import (
	"os"
	"os/signal"

	"GoTuber/CHAT/bilichat/service"
	config2 "GoTuber/CHAT/config"
)

func InitBiliChat() {
	//读取设置文件
	monitor := service.NewMonitor(config2.ChatCfg)
	monitor.Start()
	defer monitor.Stop()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill, os.Interrupt)
	<-ch
}
