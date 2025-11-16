package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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
	scanner := bufio.NewScanner(os.Stdin)

	// 保存对话历史
	messages := []aihub.ChatMessage{
		{Role: "system", Content: "你是一个友好的AI助手，请用简洁明了的方式回答问题。"},
	}

	fmt.Println("开始对话（输入 'exit' 退出）")
	fmt.Println("---")

	for {
		fmt.Print("你: ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "exit" || userInput == "退出" {
			fmt.Println("再见！")
			break
		}

		if userInput == "" {
			continue
		}

		// 添加用户消息
		messages = append(messages, aihub.ChatMessage{
			Role:    "user",
			Content: userInput,
		})

		// 发送请求
		resp, err := client.ChatCompletions(ctx, aihub.ChatCompletionRequest{
			Messages: messages,
			Temperature: floatPtr(0.7),
		})
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			continue
		}

		// 获取回复
		if len(resp.Choices) > 0 {
			assistantReply := resp.Choices[0].Message.Content
			fmt.Printf("AI: %s\n\n", assistantReply)

			// 添加助手回复到历史
			messages = append(messages, aihub.ChatMessage{
				Role:    "assistant",
				Content: assistantReply,
			})
		}
	}
}

func floatPtr(f float64) *float64 {
	return &f
}

