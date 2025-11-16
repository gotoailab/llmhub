[[中文介绍]](https://github.com/gotoailab/llmhub/blob/main/README_CN.md)  |  [[Join Discord]](https://discord.gg/B3FwBSQq)

# LLMHub - Unified LLM API Client Library

LLMHub is a unified LLM API client library developed in Golang, providing OpenAI API-compatible interfaces and supporting multiple LLM providers. It can be imported directly into your project as a library.

## Features

- ✅ **OpenAI Compatible Interface**: Fully compatible with OpenAI API format
- ✅ **Multi-Model Support**: Supports 20+ LLM providers
- ✅ **Easy to Use**: Similar usage to OpenAI SDK
- ✅ **Type Safe**: Complete type definitions
- ✅ **Streaming Response**: Supports streaming output

## Supported Model Providers

### International Models
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

### Domestic Models
- Qwen (通义千问)
- SiliconFlow (硅基流动)
- Doubao (豆包, ByteDance)
- Ernie (文心一言, Baidu)
- Spark (讯飞星火)
- ChatGLM (智谱)
- 360 Brain (360智脑)
- Hunyuan (腾讯混元, Tencent)
- Moonshot AI
- Baichuan (百川大模型)
- MINIMAX
- Yi (零一万物)
- StepFun (阶跃星辰)
- Coze

### Local Models
- Ollama

## Installation

```bash
go get github.com/gotoailab/llmhub
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/gotoailab/llmhub"
)

func main() {
    // Create client
    client, err := llmhub.NewClient(llmhub.ClientConfig{
        APIKey:   "your-api-key",
        Provider: llmhub.ProviderOpenAI,
        Model:    "gpt-3.5-turbo",
    })
    if err != nil {
        log.Fatal(err)
    }

    // Send request
    ctx := context.Background()
    resp, err := client.ChatCompletions(ctx, llmhub.ChatCompletionRequest{
        Model: "gpt-3.5-turbo",
        Messages: []llmhub.ChatMessage{
            {Role: "user", Content: "Hello, please introduce yourself"},
        },
        Temperature: floatPtr(0.7),
        MaxTokens:   intPtr(500),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Handle response
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

### Using Different Providers

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

### Streaming Response

```go
stream, err := client.ChatCompletionsStream(ctx, llmhub.ChatCompletionRequest{
    Model: "gpt-3.5-turbo",
    Messages: []llmhub.ChatMessage{
        {Role: "user", Content: "Write a poem about spring"},
    },
    Stream: true,
})
if err != nil {
    log.Fatal(err)
}
defer stream.Close()

// Read streaming data
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

## API Documentation

### Client

#### NewClient

Create a new client instance.

```go
func NewClient(config ClientConfig) (*Client, error)
```

**Parameters:**
- `config.APIKey`: API Key of the model provider (required)
- `config.Provider`: Model provider (required)
- `config.BaseURL`: Optional API base URL
- `config.Model`: Optional default model name

#### ChatCompletions

Create a chat completion request (non-streaming).

```go
func (c *Client) ChatCompletions(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionResponse, error)
```

#### ChatCompletionsStream

Create a streaming chat completion request.

```go
func (c *Client) ChatCompletionsStream(ctx context.Context, req ChatCompletionRequest) (io.ReadCloser, error)
```

#### GetProvider

Get the provider name used by the current client.

```go
func (c *Client) GetProvider() Provider
```

## Supported Provider Constants

```go
ProviderOpenAI      // OpenAI
ProviderClaude      // Anthropic Claude
ProviderDeepSeek    // DeepSeek
ProviderQwen        // Qwen (通义千问)
ProviderSiliconFlow // SiliconFlow (硅基流动)
ProviderGemini      // Google Gemini
ProviderMistral     // Mistral AI
ProviderDoubao      // Doubao (豆包)
ProviderErnie       // Ernie (文心一言)
ProviderSpark       // Spark (讯飞星火)
ProviderChatGLM     // ChatGLM
Provider360         // 360 Brain (360智脑)
ProviderHunyuan     // Hunyuan (腾讯混元)
ProviderMoonshot    // Moonshot AI
ProviderBaichuan    // Baichuan (百川)
ProviderMiniMax     // MINIMAX
ProviderGroq        // Groq
ProviderOllama      // Ollama
ProviderYi          // Yi (零一万物)
ProviderStepFun     // StepFun (阶跃星辰)
ProviderCoze        // Coze
ProviderCohere      // Cohere
ProviderTogether    // together.ai
ProviderNovita      // novita.ai
ProviderXAI         // xAI
```

## Type Definitions

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
    // ... more fields
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

## Testing

Run tests:

```bash
go test ./...
```

Run specific tests:

```bash
go test -v ./client_test.go
```

## Project Structure

```
.
├── client.go              # Client implementation
├── types.go               # Type definitions
├── provider.go            # Provider enum definitions
├── init.go                # Adapter registration
├── examples/             # Usage examples
├── internal/
│   ├── adapters/          # Model adapters
│   ├── models/            # Internal model definitions
│   └── ...
└── cmd/server/            # HTTP server (optional)
```

## Examples

See the [examples](./examples/) directory for more usage examples:

- **basic**: Basic usage example
- **multiple_providers**: Using different providers
- **stream**: Streaming response example
- **error_handling**: Error handling examples
- **conversation**: Multi-turn conversation example
- **list_providers**: List all supported providers
- **parameters**: Parameter configuration examples

## Contributing

Contributions are welcome! Please feel free to submit Issues and Pull Requests.

## License

MIT

