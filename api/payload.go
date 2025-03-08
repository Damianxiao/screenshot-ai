package api

import (
	"fmt"
)

const (
	defaultPrompt = `
现在你是一个 **专业的AI答题助理**，你的任务是 **准确、严谨、完整** 地回答我提供的图片中的题目。请你仔细阅读题目，并根据其类型，**严格遵循以下规则** 生成最优答案：

1. **选择题**（包括单选、多选）：  
   - **直接输出正确答案**，不需要解释。  
   - 如果是多选题，请提供 **所有符合题意的选项**。  
   - **答案请用加粗格式表示**（例如：` + "`**A**`、`**B, C**`" + `）。

2. **问答题**：  
   - 进行 **详细解答**，确保 **完整性、清晰性、逻辑性**。  
   - 依据 **权威知识** 进行回答，避免主观推测。  
   - **使用Markdown格式化**，答案内容可包含列表、分段、示例等，以提高可读性。  
   - **重点部分请加粗**，确保核心信息突出。  

3. **编程题**（仅限Go语言实现）：  
   - **返回完整的代码**，确保可以直接运行。  
   - **代码需包含简要注释**，解释核心逻辑，便于理解。  
   - **格式清晰，变量命名规范，遵循最佳实践**。  
   - **如涉及算法或优化策略，请补充简要解析**。

💡 **所有答案均需经过逻辑自检，确保严谨、准确，并使用Markdown格式输出！**
`
)

func NewDefaultPayload(model, base64Image string) map[string]interface{} {
	return map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []interface{}{
					map[string]string{
						"type": "text",
						"text": defaultPrompt,
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/png;base64,%s", base64Image),
						},
					},
				},
			},
		},
		"max_tokens":  3000,
		"temperature": 0.1,
		// "stream":      false,
	}
}

func NewCladue0219(model, base64Image string) map[string]interface{} {
	return map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []interface{}{
					map[string]string{
						"type": "text",
						"text": defaultPrompt,
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/png;base64,%s", base64Image),
						},
					},
				},
			},
		},
		"max_tokens":  2000,
		"temperature": 0.5,
		// "stream":      false,
	}
}

func NewGrokPayload(model, base64Image string) map[string]interface{} {
	return map[string]interface{}{

		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []interface{}{
					map[string]string{
						"type": "text",
						"text": defaultPrompt,
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/png;base64,%s", base64Image),
						},
					},
				},
			},
		},
		"max_tokens":  2000,
		"temperature": 0.5,
		// "stream":      false,
	}
}

func NewGeminiPayload(model, base64Image string) map[string]interface{} {
	return map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []interface{}{
					map[string]string{
						"type": "text",
						"text": defaultPrompt,
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/png;base64,%s", base64Image),
						},
					},
				},
			},
		},
		"max_tokens":  2000,
		"temperature": 0.5,
		// "stream":      false,
	}
}

func NewClaudePayload(model, base64Image string) map[string]interface{} {
	return map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []interface{}{
					map[string]string{
						"type": "text",
						"text": defaultPrompt,
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/png;base64,%s", base64Image),
						},
					},
				},
			},
		},
		"max_tokens":  2000,
		"temperature": 0.1,
		// "stream":      false,
	}
}

func NewGptPayload(model, base64Image string) map[string]interface{} {
	return map[string]interface{}{
		"model": model,
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []interface{}{
					map[string]string{
						"type": "text",
						"text": defaultPrompt,
					},
					map[string]interface{}{
						"type": "image_url",
						"image_url": map[string]string{
							"url": fmt.Sprintf("data:image/png;base64,%s", base64Image),
						},
					},
				},
			},
		},
		"temperature": 0.7,
		"max_tokens":  2000,
	}
}
