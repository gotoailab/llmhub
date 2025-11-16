package main

import (
	"context"
	"fmt"

	"github.com/gotoailab/llmhub"
)

func main() {
	ctx := context.Background()

	// 定义多个提供商配置
	providers := []struct {
		name     string
		provider llmhub.Provider
		apiKey   string
		model    string
	}{
		{"OpenAI", llmhub.ProviderOpenAI, "sk-your-openai-key", "gpt-3.5-turbo"},
		{"Claude", llmhub.ProviderClaude, "sk-ant-your-claude-key", "claude-3-sonnet"},
		{"DeepSeek", llmhub.ProviderDeepSeek, "your-deepseek-key", "deepseek-chat"},
		{"Qwen", llmhub.ProviderQwen, "your-qwen-key", "qwen-turbo"},
	}

	question := "用一句话解释什么是人工智能"

	for _, cfg := range providers {
		fmt.Printf("\n=== 使用 %s ===\n", cfg.name)

		client, err := llmhub.NewClient(llmhub.ClientConfig{
			APIKey:   cfg.apiKey,
			Provider: cfg.provider,
			Model:    cfg.model,
		})
		if err != nil {
			fmt.Printf("创建客户端失败: %v\n", err)
			continue
		}

		resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
			Messages: []llmhub.ChatMessage{
				{Role: "user", Content: question},
			},
		})
		if err != nil {
			fmt.Printf("请求失败: %v\n", err)
			continue
		}

		if len(resp.Choices) > 0 {
			fmt.Printf("回复: %s\n", resp.Choices[0].Message.Content)
			fmt.Printf("提供商: %s\n", client.GetProvider())
		}
	}
}
