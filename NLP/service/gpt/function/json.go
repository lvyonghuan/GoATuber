package function

import (
	"GoTuber/MEMORY"
	"GoTuber/NLP/service/gpt"
)

// InitFunctionJson 初始化JSON信息
func InitFunctionJson() {
	defer func() {
		if len(gpt.FunctionJson) == 0 {
			gpt.UseFunction = false
		} else {
			gpt.UseFunction = true
		}
	}()
	//记忆函数
	if MEMORY.MemoryCfg.IsUse {
		getMemoryJs := getMemoryJson{
			Name:        "getMemory",
			Description: "获取关于弹幕的历史记录。如果有必要可以使用。",
			Parameters: struct {
				Type       string `json:"type"`
				Properties struct {
					Chat struct {
						Type        string `json:"type"`
						Description string `json:"description"`
					} `json:"Chat"`
				} `json:"properties"`
				Required []string `json:"required"`
			}{},
		}
		getMemoryJs.Parameters.Type = "object"
		getMemoryJs.Parameters.Properties.Chat.Type = "string"
		getMemoryJs.Parameters.Properties.Chat.Description = "对user信息的提炼"
		getMemoryJs.Parameters.Required = append(getMemoryJs.Parameters.Required, "Chat")
		addFuncJson(getMemoryJs)
	}
}

//传递给OpenAI的JSON结构体

type getMemoryJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  struct {
		Type       string `json:"type"`
		Properties struct {
			Chat struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"Chat"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"parameters"`
}
