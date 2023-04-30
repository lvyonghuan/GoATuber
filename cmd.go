package main

import (
	"GoTuber/MOOD"
	"log"
	"os"
	"runtime/pprof"

	"GoTuber/CHAT"
	"GoTuber/MESSAGE"
	"GoTuber/NLP"
	"GoTuber/SPEECH"
	"GoTuber/proxy"
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
	SPEECH.InitSPEECH()
	go MOOD.InitMOOD()
	CHAT.InitChat()

	pprof.StopCPUProfile()
}
