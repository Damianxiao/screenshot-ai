package api

import (
	"fmt"
)

const (
	defaultPrompt = "请你阅读我给出的图片并理解，如果他是一道选择题，那么请你直接给出答案；如果它是一道编程题，那么请你用go语言为我实现，并返回代码。所有的回答请你用markdown格式发给我(源格式),请不要回复任何多余的话，只回复代码或答案,并且所有的回答使用中文。"
)

func NewGptPayloadWithImg(model, base64Image string) map[string]interface{} {
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
		// "max_tokens":  300,
		"max_tokens": 2000, // for test
	}
}
