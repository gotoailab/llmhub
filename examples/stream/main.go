package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
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

	// 创建流式请求
	stream, err := client.ChatCompletionsStream(ctx, aihub.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []aihub.ChatMessage{
			{Role: "user", Content: "写一首关于春天的短诗，每行不超过10个字"},
		},
		Stream: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	fmt.Println("流式响应:")
	fmt.Println("---")

	// 读取流式数据
	reader := bufio.NewReader(stream)
	buffer := make([]byte, 4096)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("读取错误: %v", err)
			break
		}

		if n > 0 {
			// 打印接收到的数据（实际使用时需要解析 SSE 格式）
			fmt.Print(string(buffer[:n]))
		}
	}

	fmt.Println("\n---")
	fmt.Println("流式响应完成")
}

