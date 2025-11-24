package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gotoailab/llmhub"
)

func main() {
	// 创建支持工具调用的客户端（OpenAI）
	client, err := llmhub.NewClient(llmhub.ClientConfig{
		APIKey:   "your-openai-api-key",
		Provider: llmhub.ProviderOpenAI,
		Model:    "gpt-4",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 定义一个简单的工具：获取天气信息
	weatherTool := llmhub.Tool{
		Type: "function",
		Function: llmhub.FunctionDefinition{
			Name:        "get_current_weather",
			Description: "Get the current weather in a given location",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"location": map[string]interface{}{
						"type":        "string",
						"description": "The city and state, e.g. San Francisco, CA",
					},
					"unit": map[string]interface{}{
						"type": "string",
						"enum": []string{"celsius", "fahrenheit"},
					},
				},
				"required": []string{"location"},
			},
		},
	}

	// 发送带工具的聊天请求
	ctx := context.Background()
	resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
		Model: "gpt-4",
		Messages: []llmhub.ChatMessage{
			{Role: "user", Content: "What's the weather like in Boston?"},
		},
		Tools: []llmhub.Tool{weatherTool},
		ToolChoice: "auto",
		Temperature: floatPtr(0.7),
		MaxTokens:   intPtr(500),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 处理响应
	if len(resp.Choices) > 0 {
		message := resp.Choices[0].Message
		
		// 检查是否有工具调用
		if len(message.ToolCalls) > 0 {
			fmt.Println("模型请求调用工具:")
			for _, toolCall := range message.ToolCalls {
				fmt.Printf("工具: %s\n", toolCall.Function.Name)
				fmt.Printf("参数: %s\n", toolCall.Function.Arguments)
				
				// 模拟工具执行
				toolResult := executeWeatherTool(toolCall.Function.Arguments)
				
				// 将工具结果发送回模型
				finalResp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
					Model: "gpt-4",
					Messages: []llmhub.ChatMessage{
						{Role: "user", Content: "What's the weather like in Boston?"},
						{Role: "assistant", Content: message.Content, ToolCalls: message.ToolCalls},
						{Role: "tool", Content: toolResult, ToolCallID: toolCall.ID},
					},
					Tools: []llmhub.Tool{weatherTool},
				})
				if err != nil {
					log.Fatal(err)
				}
				
				if len(finalResp.Choices) > 0 {
					fmt.Println("最终回复:", finalResp.Choices[0].Message.Content)
				}
			}
		} else {
			fmt.Println("直接回复:", message.Content)
		}
	}

	// 测试不支持工具调用的模型
	fmt.Println("\n测试不支持工具调用的模型:")
	qwenClient, err := llmhub.NewClient(llmhub.ClientConfig{
		APIKey:   "your-qwen-api-key",
		Provider: llmhub.ProviderQwen,
		Model:    "qwen-turbo",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 尝试使用工具调用（应该会失败）
	_, err = qwenClient.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
		Model: "qwen-turbo",
		Messages: []llmhub.ChatMessage{
			{Role: "user", Content: "What's the weather like in Boston?"},
		},
		Tools: []llmhub.Tool{weatherTool},
	})
	if err != nil {
		fmt.Printf("期望的错误: %v\n", err)
	}
}

// 模拟天气工具的执行
func executeWeatherTool(arguments string) string {
	var params map[string]interface{}
	json.Unmarshal([]byte(arguments), &params)
	
	location := params["location"].(string)
	unit := "celsius"
	if u, ok := params["unit"].(string); ok {
		unit = u
	}
	
	// 模拟天气数据
	weather := map[string]interface{}{
		"location":    location,
		"temperature": 22,
		"unit":        unit,
		"description": "Sunny",
	}
	
	result, _ := json.Marshal(weather)
	return string(result)
}

func floatPtr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}