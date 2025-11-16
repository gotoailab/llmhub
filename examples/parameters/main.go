package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aihub"
)

func main() {
	client, err := aihub.NewClient(aihub.ClientConfig{
		APIKey:   "your-api-key-here",
		Provider: aihub.ProviderOpenAI,
		Model:    "gpt-3.5-turbo",
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	question := "解释一下量子计算的基本原理"

	// 示例1: 使用默认参数
	fmt.Println("=== 示例1: 默认参数 ===")
	resp1, _ := client.ChatCompletions(ctx, aihub.ChatCompletionRequest{
		Messages: []aihub.ChatMessage{
			{Role: "user", Content: question},
		},
	})
	if len(resp1.Choices) > 0 {
		fmt.Printf("回复长度: %d 字符\n", len(resp1.Choices[0].Message.Content.(string)))
	}
	fmt.Println()

	// 示例2: 低温度（更确定性）
	fmt.Println("=== 示例2: 低温度 (Temperature=0.3) ===")
	resp2, _ := client.ChatCompletions(ctx, aihub.ChatCompletionRequest{
		Messages: []aihub.ChatMessage{
			{Role: "user", Content: question},
		},
		Temperature: floatPtr(0.3),
	})
	if len(resp2.Choices) > 0 {
		fmt.Printf("回复长度: %d 字符\n", len(resp2.Choices[0].Message.Content.(string)))
	}
	fmt.Println()

	// 示例3: 高温度（更创造性）
	fmt.Println("=== 示例3: 高温度 (Temperature=1.0) ===")
	resp3, _ := client.ChatCompletions(ctx, aihub.ChatCompletionRequest{
		Messages: []aihub.ChatMessage{
			{Role: "user", Content: question},
		},
		Temperature: floatPtr(1.0),
	})
	if len(resp3.Choices) > 0 {
		fmt.Printf("回复长度: %d 字符\n", len(resp3.Choices[0].Message.Content.(string)))
	}
	fmt.Println()

	// 示例4: 限制最大 Token 数
	fmt.Println("=== 示例4: 限制最大 Token (MaxTokens=50) ===")
	resp4, _ := client.ChatCompletions(ctx, aihub.ChatCompletionRequest{
		Messages: []aihub.ChatMessage{
			{Role: "user", Content: question},
		},
		MaxTokens: intPtr(50),
	})
	if len(resp4.Choices) > 0 {
		fmt.Printf("回复: %s\n", resp4.Choices[0].Message.Content)
		fmt.Printf("Token 使用: %d\n", resp4.Usage.TotalTokens)
	}
	fmt.Println()

	// 示例5: 使用停止词
	fmt.Println("=== 示例5: 使用停止词 ===")
	resp5, _ := client.ChatCompletions(ctx, aihub.ChatCompletionRequest{
		Messages: []aihub.ChatMessage{
			{Role: "user", Content: "列出三个编程语言，每个一行"},
		},
		Stop: []string{"\n\n"}, // 遇到两个换行符时停止
	})
	if len(resp5.Choices) > 0 {
		fmt.Printf("回复: %s\n", resp5.Choices[0].Message.Content)
	}
}

func floatPtr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}

