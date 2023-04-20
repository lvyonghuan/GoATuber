package gpt

const Openaiapiurl1 = "https://api.openai.com/v1/chat/completions" //对话使用的url

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 对话使用的Request body
type postData struct {
	Model            string             `json:"model"`
	Messages         []Messages         `json:"messages"` //message依靠传入信息获取
	Role             string             `json:"role"`     //角色信息依靠获取用户名来完成，和message一起传入NLP模块。
	MaxTokens        int                `json:"max_tokens"`
	Temperature      float64            `json:"temperature"`
	TopP             float32            `json:"top_p"`
	Stop             []string           `json:"stop"`
	PresencePenalty  float32            `json:"presence_penalty"`
	FrequencyPenalty float32            `json:"frequency_penalty"`
	LogitBias        map[string]float32 `json:"logit_bias"`
}

// OpenAiRcv 对话使用的Response
type OpenAiRcv struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message      Messages `json:"message"`
		FinishReason string   `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens    int `json:"prompt_tokens"`
		CompletionTokes int `json:"completion_tokens"`
		TotalTokens     int `json:"total_tokens"`
	}
}

func GenerateText() {

}
