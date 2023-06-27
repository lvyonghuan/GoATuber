package function

import memory_gpt "GoTuber/MEMORY/NLPmodel/gpt"

//如果要做更多功能，比如搜索网站什么的，这里就容得下的入口。我建议把这里当入口，或者实现一些代码量比较小的函数。
//函数统一：func([]string) string。建议说明切片的每个索引所代表的含义。

// 获取记忆。索引0表示信息。
func getMemory(ms []string) string {
	//调用memory模块实现记忆查询
	chat := memory_gpt.Chat{
		Human: ms[0],
	}
	human, ai, user := chat.GetMemory()
	return "user:" + user + "说：" + human + "assistant:" + ai
}
