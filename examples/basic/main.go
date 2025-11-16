package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aihub"
)

func main() {
	// 创建客户端
	client, err := aihub.NewClient(aihub.ClientConfig{
		APIKey:   "your-api-key-here",
		Provider: aihub.ProviderOpenAI,
		Model:    "gpt-3.5-turbo",
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// 发送聊天请求
	resp, err := client.ChatCompletions(ctx, aihub.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []aihub.ChatMessage{
			{Role: "user", Content: "你好，请用一句话介绍一下你自己"},
		},
		Temperature: floatPtr(0.7),
		MaxTokens:   intPtr(100),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 处理响应
	if len(resp.Choices) > 0 {
		fmt.Println("回复:", resp.Choices[0].Message.Content)
		fmt.Printf("Token 使用: %d (输入: %d, 输出: %d)\n",
			resp.Usage.TotalTokens,
			resp.Usage.PromptTokens,
			resp.Usage.CompletionTokens)
	} else {
		log.Fatal("没有回复")
	}
}

func floatPtr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}
