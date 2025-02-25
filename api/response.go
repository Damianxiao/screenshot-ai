package api

// GptResponse 定义 OpenAI Chat Completions API 的响应结构
type GptResponse struct {
	ID      string          `json:"id"`
	Created int64           `json:"created"`
	Model   string          `json:"model"`
	Choices []GptRespChoice `json:"choices"`
	Usage   GptRespUsage    `json:"usage"`
}

type GptRespChoice struct {
	Message GptRespMessage `json:"message"`
}

type GptRespMessage struct {
	Content string `json:"content"`
}

type GptRespUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
