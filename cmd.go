package main

import (
	"GoTuber/CHAT"
	"GoTuber/MEMORY"
	"GoTuber/MESSAGE"
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MOOD"
	"GoTuber/NLP"
	"GoTuber/SPEECH"
	"GoTuber/VOICE"
	"GoTuber/frontend/live2d_backend"
	"GoTuber/proxy"
	"log"
)

func main() {
	//f, err := os.Create("profile.pb")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	log.Fatal(err)
	//}
	//
	log.Println("Go!")
	sensitive.InitConfig()
	go MESSAGE.GetMessage()
	proxy.InitProxyConfig()
	MEMORY.InitMemory()
	NLP.InitNLP()
	VOICE.InitVOICE()
	SPEECH.InitConfig()
	go MOOD.InitMOOD()
	go backend.Init()
	go MESSAGE.GetMessage()
	CHAT.InitChat()
	//
	//pprof.StopCPUProfile()
}
