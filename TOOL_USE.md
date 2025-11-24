# Tool Use 功能文档

LLMHub 现在支持大模型的工具调用（Tool Use / Function Calling）功能。

## 支持的模型厂商

### ✅ 完全支持工具调用的厂商
- **OpenAI**: GPT-4, GPT-4 Turbo, GPT-3.5 Turbo
- **Anthropic Claude**: Claude-3.5 Sonnet, Claude-3 Opus, Claude-3 Sonnet
- **OpenRouter**: 支持的模型
- **Groq**: 支持的模型
- **DeepSeek**: 支持的模型
- **Together AI**: 支持的模型
- **SiliconFlow**: 支持的模型
- **Moonshot**: 支持的模型
- **StepFun**: 支持的模型
- **Mistral**: 支持的模型
- **Cohere**: 支持的模型
- **Novita**: 支持的模型
- **xAI**: 支持的模型

### ❌ 暂不支持工具调用的厂商
- **通义千问 (Qwen)**: 使用自有格式，暂未集成
- **百川 (Baichuan)**
- **智谱 GLM (ChatGLM)**
- **文心一言 (Ernie)**
- **讯飞星火 (Spark)**
- **腾讯混元 (Hunyuan)**
- **360智脑**
- **MiniMax**
- **零一万物 (Yi)**
- **豆包 (Doubao)**
- **Ollama**: 需要特殊处理
- **Coze**

## 使用示例

### 基本工具调用

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    
    "github.com/gotoailab/llmhub"
)

func main() {
    // 创建客户端
    client, err := llmhub.NewClient(llmhub.ClientConfig{
        APIKey:   "your-api-key",
        Provider: llmhub.ProviderOpenAI,
        Model:    "gpt-4",
    })
    if err != nil {
        log.Fatal(err)
    }

    // 定义工具
    weatherTool := llmhub.Tool{
        Type: "function",
        Function: llmhub.FunctionDefinition{
            Name:        "get_weather",
            Description: "获取指定城市的天气信息",
            Parameters: map[string]interface{}{
                "type": "object",
                "properties": map[string]interface{}{
                    "city": map[string]interface{}{
                        "type":        "string",
                        "description": "城市名称",
                    },
                    "unit": map[string]interface{}{
                        "type": "string",
                        "enum": []string{"celsius", "fahrenheit"},
                    },
                },
                "required": []string{"city"},
            },
        },
    }

    // 发送请求
    resp, err := client.ChatCompletions(context.Background(), llmhub.ChatCompletionRequest{
        Model: "gpt-4",
        Messages: []llmhub.ChatMessage{
            {Role: "user", Content: "北京今天天气怎么样？"},
        },
        Tools:       []llmhub.Tool{weatherTool},
        ToolChoice:  "auto",
        Temperature: &[]float64{0.7}[0],
    })
    if err != nil {
        log.Fatal(err)
    }

    // 处理工具调用
    if len(resp.Choices) > 0 && len(resp.Choices[0].Message.ToolCalls) > 0 {
        for _, toolCall := range resp.Choices[0].Message.ToolCalls {
            fmt.Printf("工具调用: %s\n", toolCall.Function.Name)
            fmt.Printf("参数: %s\n", toolCall.Function.Arguments)
            
            // 执行工具并返回结果
            result := executeWeatherTool(toolCall.Function.Arguments)
            
            // 发送工具结果
            finalResp, err := client.ChatCompletions(context.Background(), llmhub.ChatCompletionRequest{
                Model: "gpt-4",
                Messages: []llmhub.ChatMessage{
                    {Role: "user", Content: "北京今天天气怎么样？"},
                    {Role: "assistant", ToolCalls: resp.Choices[0].Message.ToolCalls},
                    {Role: "tool", Content: result, ToolCallID: toolCall.ID},
                },
                Tools: []llmhub.Tool{weatherTool},
            })
            if err != nil {
                log.Fatal(err)
            }
            
            fmt.Println("最终回复:", finalResp.Choices[0].Message.Content)
        }
    }
}

func executeWeatherTool(arguments string) string {
    var params map[string]interface{}
    json.Unmarshal([]byte(arguments), &params)
    
    // 模拟天气数据
    weather := map[string]interface{}{
        "city":        params["city"],
        "temperature": 22,
        "unit":        "celsius",
        "description": "晴朗",
    }
    
    result, _ := json.Marshal(weather)
    return string(result)
}
```

### 多个工具调用

```go
// 定义多个工具
tools := []llmhub.Tool{
    {
        Type: "function",
        Function: llmhub.FunctionDefinition{
            Name:        "get_weather",
            Description: "获取天气信息",
            Parameters: /* ... */,
        },
    },
    {
        Type: "function",
        Function: llmhub.FunctionDefinition{
            Name:        "search_news",
            Description: "搜索新闻",
            Parameters: /* ... */,
        },
    },
}

// 使用多个工具
resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
    Model:    "gpt-4",
    Messages: messages,
    Tools:    tools,
    ToolChoice: "auto", // 让模型自动选择
})
```

### 强制使用特定工具

```go
// 强制使用特定工具
resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
    Model:    "gpt-4",
    Messages: messages,
    Tools:    tools,
    ToolChoice: map[string]interface{}{
        "type": "function",
        "function": map[string]string{
            "name": "get_weather",
        },
    },
})
```

### 禁用工具调用

```go
// 禁用工具调用，即使定义了工具
resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
    Model:      "gpt-4",
    Messages:   messages,
    Tools:      tools,
    ToolChoice: "none", // 禁用工具调用
})
```

## 错误处理

当模型不支持工具调用时，会返回相应的错误：

```go
resp, err := qwenClient.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
    Model: "qwen-turbo",
    Messages: []llmhub.ChatMessage{
        {Role: "user", Content: "帮我查一下天气"},
    },
    Tools: tools, // Qwen 不支持工具调用
})
if err != nil {
    // 错误: "tool use not supported for Qwen models"
    fmt.Printf("错误: %v\n", err)
}
```

## 注意事项

1. **工具参数格式**: 工具参数必须是 JSON Schema 格式
2. **工具结果**: 工具执行结果需要以字符串形式返回
3. **消息顺序**: 工具调用的消息顺序为：用户消息 → 助手工具调用 → 工具结果 → 助手最终回复
4. **流式响应**: 工具调用同样支持流式响应
5. **兼容性**: 同时支持新版 `tools` 格式和旧版 `functions` 格式

## 支持的工具选择策略

- `"auto"`: 让模型自动决定是否调用工具
- `"none"`: 禁用工具调用
- `{"type": "function", "function": {"name": "tool_name"}}`: 强制调用指定工具

## Claude 特殊处理

Claude 模型有一些特殊要求：
- Claude-3 Haiku 不支持工具调用
- System 消息会被特殊处理
- 工具结果格式略有不同

## 完整的工具调用流程

1. 定义工具和参数 Schema
2. 发送包含工具的聊天请求
3. 检查响应中是否有工具调用
4. 执行相应的工具函数
5. 将工具结果发送回模型
6. 获取模型的最终回复