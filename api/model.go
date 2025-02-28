package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"screenshot-ai/config"
	"screenshot-ai/pkg/log"
)

func CallApiWithimg(img, model string) (string, error) {
	// TODO set pattern
	modelConfig := config.Cfg.Model[model[:]]
	fmt.Println("call ai.....")
	payload := selectModel(model, img)
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("api:JSON marshal failed: %v", err)
	}
	log.Logobj.Info("call model: ", model)
	msg, err := Call(modelConfig.Url, modelConfig.Api, jsonPayload)
	// broadcast to hub
	if err != nil {
		return "", err
	}
	return msg, nil
}

func selectModel(model, img string) map[string]interface{} {
	switch model {
	case "gpt-4o":
		// 为gpt-4o模型执行特定操作
		return NewGptPayload(model, img)
	case "claude-3-5-sonnet-20240620":
		// 为another-model模型执行另一特定操作
		return NewClaudePayload(model, img)
	case "gemini-1.5-flash":
		return NewGeminiPayload(model, img)
	case "grok-2-vision-1212":
		return NewGrokPayload(model, img)
	case "claude-3-7-sonnet-20250219":
		return NewCladue0219(model,img)		
	default:
		
		// 为其他所有模型执行默认操作
		return NewGptPayload(model, img)
	}
}

func Call(url, apiKey string, jsonPayload []byte) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("call:create req failed")
	}
	// set api
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("call:request ai failed,%s", err)

	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("call:read resp body failed")
	}
	fmt.Println(string(data))
	var res GptResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return "", fmt.Errorf("call:unmarshal resp body failed")
	}
	if len(res.Choices) == 0 {
		return "", fmt.Errorf("resp:invalid resp")
	}
	log.Logobj.Info("model resp:", res.Choices[0].Message.Content)
	msg := res.Choices[0].Message.Content

	return msg, nil
}
