package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gotoailab/llmhub"
)

func main() {
	ctx := context.Background()

	// 示例1: 无效的 Provider
	fmt.Println("=== 示例1: 无效的 Provider ===")
	_, err := llmhub.NewClient(llmhub.ClientConfig{
		APIKey:   "test-key",
		Provider: llmhub.Provider("invalid-provider"),
	})
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
	fmt.Println()

	// 示例2: 缺少 API Key
	fmt.Println("=== 示例2: 缺少 API Key ===")
	_, err = llmhub.NewClient(llmhub.ClientConfig{
		Provider: llmhub.ProviderOpenAI,
	})
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
	fmt.Println()

	// 示例3: 缺少模型名称
	fmt.Println("=== 示例3: 缺少模型名称 ===")
	client, err := llmhub.NewClient(llmhub.ClientConfig{
		APIKey:   "test-key",
		Provider: llmhub.ProviderOpenAI,
		// 注意：没有设置 Model
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
		// 请求中也没有指定 Model
		Messages: []llmhub.ChatMessage{
			{Role: "user", Content: "Hello"},
		},
	})
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}
	fmt.Println()

	// 示例4: 无效的 API Key（实际调用会失败）
	fmt.Println("=== 示例4: 无效的 API Key ===")
	client, err = llmhub.NewClient(llmhub.ClientConfig{
		APIKey:   "invalid-key",
		Provider: llmhub.ProviderOpenAI,
		Model:    "gpt-3.5-turbo",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
		Messages: []llmhub.ChatMessage{
			{Role: "user", Content: "Hello"},
		},
	})
	if err != nil {
		fmt.Printf("API 调用错误: %v\n", err)
	}
}
