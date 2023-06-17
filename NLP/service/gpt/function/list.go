package function

import "GoTuber/MEMORY"

func InitFunction() {
	//如果启用记忆模式，则启用记忆调取函数
	if MEMORY.MemoryCfg.IsUse {
		addFunc(getMemory)
	}
}
