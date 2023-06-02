package main

import (
	"GoTuber/CHAT"
	"GoTuber/MEMORY"
	"GoTuber/MESSAGE"
	"GoTuber/MOOD"
	"GoTuber/NLP"
	"GoTuber/SPEECH"
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
	go MESSAGE.GetMessage()
	proxy.InitProxyConfig()
	NLP.InitNLP()
	MEMORY.InitMemory()
	SPEECH.InitSPEECH()
	go MOOD.InitMOOD()
	go backend.Init()
	CHAT.InitChat()
	//
	//pprof.StopCPUProfile()
}
