package main

import (
	"log"

	"GoTuber/CHAT"
	"GoTuber/MEMORY"
	"GoTuber/MESSAGE"
	sensitive "GoTuber/MESSAGE/filter"
	"GoTuber/MOOD"
	"GoTuber/NLP"
	"GoTuber/SPEECH"
	"GoTuber/frontend/model_backend"
	"GoTuber/proxy"
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
	SPEECH.InitSPEECH()
	go MOOD.InitMOOD()
	go backend.Init()
	CHAT.InitChat()
	//
	//pprof.StopCPUProfile()
}
