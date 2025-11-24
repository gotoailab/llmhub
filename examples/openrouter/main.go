package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gotoailab/llmhub"
)

func main() {
	// 创建 OpenRouter 客户端
	client, err := llmhub.NewClient(llmhub.ClientConfig{
		APIKey:   "your-openrouter-api-key",
		Provider: llmhub.ProviderOpenRouter,
		Model:    "openai/gpt-3.5-turbo", // OpenRouter 使用 provider/model 格式
	})
	if err != nil {
		log.Fatal(err)
	}

	// 发送聊天请求
	ctx := context.Background()
	resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
		Model: "openai/gpt-3.5-turbo",
		Messages: []llmhub.ChatMessage{
			{Role: "user", Content: "Hello, can you introduce yourself?"},
		},
		Temperature: floatPtr(0.7),
		MaxTokens:   intPtr(500),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 输出响应
	if len(resp.Choices) > 0 {
		fmt.Println("OpenRouter Response:", resp.Choices[0].Message.Content)
	}

	// 使用其他模型的示例
	resp2, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
		Model: "anthropic/claude-3-haiku",
		Messages: []llmhub.ChatMessage{
			{Role: "user", Content: "What is OpenRouter?"},
		},
		Temperature: floatPtr(0.7),
		MaxTokens:   intPtr(300),
	})
	if err != nil {
		log.Fatal(err)
	}

	if len(resp2.Choices) > 0 {
		fmt.Println("Claude via OpenRouter:", resp2.Choices[0].Message.Content)
	}
}

func floatPtr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}