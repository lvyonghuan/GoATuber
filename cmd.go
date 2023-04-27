package main

import (
	"GoTuber/CHAT"
	"GoTuber/MESSAGE"
	"GoTuber/NLP"
	"GoTuber/proxy"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("profile.pb")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}

	log.Println("Go!")
	go MESSAGE.GetMessage()
	proxy.InitProxyConfig()
	NLP.InitNLP()
	CHAT.InitChat()

	pprof.StopCPUProfile()
}
