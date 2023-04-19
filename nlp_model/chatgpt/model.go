package chatgpt

const Openaiapiurl1 = "https://api.openai.com/v1/chat/completions" //对话使用的url
const Openaiapiurl2 = "https://api.openai.com/v1/completions"      //角色扮演使用的url

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 对话使用的Request body
type postData struct {
	Model       string     `json:"model"`
	Messages    []Messages `json:"messages"`
	MaxTokens   int        `json:"max_tokens"`
	Temperature float64    `json:"temperature"`
}

// 角色扮演使用的Request body
type postDataWithIdentity struct {
	Model       string        `json:"model"`
	MaxTokens   int           `json:"max_tokens"`
	Temperature float64       `json:"temperature"`
	Prompt      []interface{} `json:"prompt"`
	Stop        []string      `json:"stop"`
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

// OpenAiRcvWithIdentity 角色扮演使用的Response
type OpenAiRcvWithIdentity struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     int    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens    int `json:"prompt_tokens"`
		CompletionTokes int `json:"completion_tokens"`
		TotalTokens     int `json:"total_tokens"`
	}
}
