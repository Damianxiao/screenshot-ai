package api

import (
	"fmt"
)

const (
	defaultPrompt = "请你阅读我给出的图片并理解，如果他是一道选择题，那么请你直接给出答案(可能是多选，给出全部符合题意的选项);如果它是一道问答题,请你根据题意尽可能详细的给出回答；如果它是一道编程题;那么请你用go语言为我实现，并返回代码。所有的回答请你用markdown格式发给我(源格式)"
)

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
		"temperature": 0.5,
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
