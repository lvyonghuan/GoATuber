package MOOD

import dictMOOD "GoTuber/MOOD/dict"

func InitMOOD() {
	go dictMOOD.Search()
}
