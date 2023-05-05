package MOOD

import dictMOOD "GoTuber/MOOD/dict"

func InitMOOD() {
	//go readMoodAct()
	go dictMOOD.Search()
}
