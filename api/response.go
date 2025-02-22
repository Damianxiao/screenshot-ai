package api

// GptResponse 定义 OpenAI Chat Completions API 的响应结构
type GptResponse struct {
	ID      string          `json:"id"`      // 请求 ID，用于日志追踪
	Created int64           `json:"created"` // 创建时间，用于日志
	Model   string          `json:"model"`   // 模型版本，用于日志
	Choices []GptRespChoice `json:"choices"` // 包含回复内容
	Usage   GptRespUsage    `json:"usage"`   // token 使用情况，用于日志
}

type GptRespChoice struct {
	Message GptRespMessage `json:"message"` // 回复消息
}

type GptRespMessage struct {
	Content string `json:"content"` // GPT 的回答
}

type GptRespUsage struct {
	PromptTokens     int `json:"prompt_tokens"`     // 输入 token
	CompletionTokens int `json:"completion_tokens"` // 输出 token
	TotalTokens      int `json:"total_tokens"`      // 总 token
}
