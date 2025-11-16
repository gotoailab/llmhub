# LLMHub - 统一大模型 API 客户端库

LLMHub 是一个用 Golang 开发的统一大模型 API 客户端库，提供与 OpenAI API 对齐的接口，支持多种大模型提供商。可以作为库直接导入到你的项目中。

## 功能特性

- ✅ **OpenAI 兼容接口**：完全兼容 OpenAI API 格式
- ✅ **多模型支持**：支持 20+ 种大模型提供商
- ✅ **简单易用**：类似 OpenAI SDK 的使用方式
- ✅ **类型安全**：完整的类型定义
- ✅ **流式响应**：支持流式输出

## 支持的模型提供商

### 国际模型
- OpenAI (GPT-3.5, GPT-4, etc.)
- Anthropic Claude
- Google Gemini / PaLM2
- Mistral AI
- DeepSeek
- Groq
- Cohere
- xAI
- Together.ai
- Novita.ai

### 国内模型
- Qwen (通义千问)
- 硅基流动 (SiliconFlow)
- 豆包（字节跳动）
- 文心一言（百度）
- 讯飞星火
- ChatGLM（智谱）
- 360智脑
- 腾讯混元
- Moonshot AI
- 百川大模型
- MINIMAX
- 零一万物
- 阶跃星辰
- Coze

### 本地模型
- Ollama

## 安装

```bash
go get github.com/gotoailab/llmhub
```

## 快速开始

### 基本使用

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/gotoailab/llmhub"
)

func main() {
    // 创建客户端
    client, err := llmhub.NewClient(llmhub.ClientConfig{
        APIKey:   "your-api-key",
        Provider: llmhub.ProviderOpenAI,
        Model:    "gpt-3.5-turbo",
    })
    if err != nil {
        log.Fatal(err)
    }

    // 发送请求
    ctx := context.Background()
    resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
        Model: "gpt-3.5-turbo",
        Messages: []llmhub.ChatMessage{
            {Role: "user", Content: "你好，请介绍一下你自己"},
        },
        Temperature: floatPtr(0.7),
        MaxTokens:   intPtr(500),
    })
    if err != nil {
        log.Fatal(err)
    }

    // 处理响应
    if len(resp.Choices) > 0 {
        fmt.Println(resp.Choices[0].Message.Content)
    }
}

func floatPtr(f float64) *float64 {
    return &f
}

func intPtr(i int) *int {
    return &i
}
```

### 使用不同提供商

```go
// Claude
client, _ := llmhub.NewClient(llmhub.ClientConfig{
    APIKey:   "sk-ant-your-key",
    Provider: llmhub.ProviderClaude,
    Model:    "claude-3-sonnet",
})

// DeepSeek
client, _ := llmhub.NewClient(llmhub.ClientConfig{
    APIKey:   "your-deepseek-key",
    Provider: llmhub.ProviderDeepSeek,
    Model:    "deepseek-chat",
})

// Qwen
client, _ := llmhub.NewClient(llmhub.ClientConfig{
    APIKey:   "your-qwen-key",
    Provider: llmhub.ProviderQwen,
    Model:    "qwen-turbo",
})
```

### 流式响应

```go
stream, err := client.ChatCompletionsStream(ctx, llmhub.ChatCompletionRequest{
    Model: "gpt-3.5-turbo",
    Messages: []llmhub.ChatMessage{
        {Role: "user", Content: "写一首关于春天的诗"},
    },
    Stream: true,
})
if err != nil {
    log.Fatal(err)
}
defer stream.Close()

// 读取流式数据
buf := make([]byte, 4096)
for {
    n, err := stream.Read(buf)
    if err != nil {
        if err == io.EOF {
            break
        }
        log.Fatal(err)
    }
    fmt.Print(string(buf[:n]))
}
```

## API 文档

### Client

#### NewClient

创建新的客户端实例。

```go
func NewClient(config ClientConfig) (*Client, error)
```

**参数：**
- `config.APIKey`: 模型提供商的 API Key（必需）
- `config.Provider`: 模型提供商（必需）
- `config.BaseURL`: 可选的 API 基础 URL
- `config.Model`: 可选的默认模型名称

#### ChatCompletions

创建聊天完成请求（非流式）。

```go
func (c *Client) ChatCompletions(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionResponse, error)
```

#### ChatCompletionsStream

创建流式聊天完成请求。

```go
func (c *Client) ChatCompletionsStream(ctx context.Context, req ChatCompletionRequest) (io.ReadCloser, error)
```

#### GetProvider

获取当前客户端使用的提供商名称。

```go
func (c *Client) GetProvider() string
```

## 支持的 Provider 常量

```go
ProviderOpenAI      // OpenAI
ProviderClaude      // Anthropic Claude
ProviderDeepSeek    // DeepSeek
ProviderQwen        // 通义千问
ProviderSiliconFlow // 硅基流动
ProviderGemini      // Google Gemini
ProviderMistral     // Mistral AI
ProviderDoubao      // 豆包
ProviderErnie       // 文心一言
ProviderSpark       // 讯飞星火
ProviderChatGLM     // ChatGLM
Provider360         // 360智脑
ProviderHunyuan     // 腾讯混元
ProviderMoonshot    // Moonshot AI
ProviderBaichuan    // 百川
ProviderMiniMax     // MINIMAX
ProviderGroq        // Groq
ProviderOllama      // Ollama
ProviderYi          // 零一万物
ProviderStepFun     // 阶跃星辰
ProviderCoze        // Coze
ProviderCohere      // Cohere
ProviderTogether    // together.ai
ProviderNovita      // novita.ai
ProviderXAI         // xAI
```

## 类型定义

### ChatCompletionRequest

```go
type ChatCompletionRequest struct {
    Model            string
    Messages         []ChatMessage
    Temperature      *float64
    TopP             *float64
    MaxTokens        *int
    Stream           bool
    PresencePenalty  *float64
    FrequencyPenalty *float64
    Stop             []string
    // ... 更多字段
}
```

### ChatCompletionResponse

```go
type ChatCompletionResponse struct {
    ID      string
    Object  string
    Created int64
    Model  string
    Choices []ChatCompletionChoice
    Usage   Usage
}
```

## 测试

运行测试：

```bash
go test ./...
```

运行特定测试：

```bash
go test -v ./client_test.go
```

## 项目结构

```
.
├── client.go              # 客户端实现
├── types.go               # 类型定义
├── init.go                # 适配器注册
├── internal/
│   ├── adapters/          # 模型适配器
│   ├── models/            # 内部模型定义
│   └── ...
└── cmd/server/            # HTTP 服务（可选）
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## License

MIT
