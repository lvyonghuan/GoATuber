package gpt

var MS []Messages     //向OpenAI传递的消息，包含了用户设定的提示词
var roleMS []Messages //角色信息

const OpenAIChatUrl = "https://api.openai.com/v1/chat/completions" //OpenAI对话使用的url

type RequestMessages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ReceiveMessages struct {
	Role         string `json:"role"`
	Content      string `json:"content"`
	FunctionCall struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function_call"`
}

// 对话使用的Request body
type postData struct {
	Model            string            `json:"model"`
	Messages         []RequestMessages `json:"messages"` //message依靠传入信息获取
	MaxTokens        int               `json:"max_tokens"`
	Temperature      float64           `json:"temperature"`
	TopP             float64           `json:"top_p"`
	Stop             string            `json:"stop"`
	PresencePenalty  float64           `json:"presence_penalty"`
	FrequencyPenalty float64           `json:"frequency_penalty"`
	User             string            `json:"user"`
}

// OpenAiRcv Response
type OpenAiRcv struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message      ReceiveMessages `json:"message"`
		FinishReason string          `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens    int `json:"prompt_tokens"`
		CompletionTokes int `json:"completion_tokens"`
		TotalTokens     int `json:"total_tokens"`
	}
}
