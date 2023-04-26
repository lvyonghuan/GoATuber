package main

import (
	"GoTuber/CHAT"
	"GoTuber/MESSAGE"
	"GoTuber/NLP"
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

	go MESSAGE.GetMessage()
	NLP.InitNLP()
	CHAT.InitChat()

	pprof.StopCPUProfile()
}
