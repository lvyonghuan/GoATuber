package MOOD

import (
	sensitive "GoTuber/MESSAGE/filter"
	dictMOOD "GoTuber/MOOD/dict"
)

func GetMessage(msg *sensitive.OutPut) {
	dictMOOD.HandelMsg(msg)
}
