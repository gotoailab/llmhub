package adapters

import (
	"github.com/aihub/internal/models"
	"context"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
)

type OpenAIAdapter struct {
	client *openai.Client
}

// NewOpenAIAdapter 创建 OpenAI 适配器（导出以供注册）
func NewOpenAIAdapter(apiKey, baseURL string) (Adapter, error) {
	config := openai.DefaultConfig(apiKey)
	if baseURL != "" {
		config.BaseURL = baseURL
	}
	client := openai.NewClientWithConfig(config)
	
	return &OpenAIAdapter{
		client: client,
	}, nil
}

func (a *OpenAIAdapter) GetProvider() Provider {
	return Provider("openai")
}

func (a *OpenAIAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	// 转换消息格式
	messages := make([]openai.ChatCompletionMessage, 0, len(req.Messages))
	for _, msg := range req.Messages {
		var content string
		switch v := msg.Content.(type) {
		case string:
			content = v
		case []interface{}:
			// 处理多模态内容
			content = fmt.Sprintf("%v", v)
		default:
			content = fmt.Sprintf("%v", v)
		}
		
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: content,
		})
	}

	// 构建请求
	openaiReq := openai.ChatCompletionRequest{
		Model:       req.Model,
		Messages:    messages,
		Temperature: float32(getFloatValue(req.Temperature)),
		TopP:        float32(getFloatValue(req.TopP)),
		MaxTokens:   getIntValue(req.MaxTokens),
		Stream:      false,
	}

	if req.Stop != nil && len(req.Stop) > 0 {
		openaiReq.Stop = req.Stop
	}

	// 调用 OpenAI API
	resp, err := a.client.CreateChatCompletion(ctx, openaiReq)
	if err != nil {
		return nil, fmt.Errorf("openai api error: %w", err)
	}

	// 转换响应格式
	choices := make([]models.ChatCompletionChoice, 0, len(resp.Choices))
	for _, choice := range resp.Choices {
		choices = append(choices, models.ChatCompletionChoice{
			Index: choice.Index,
			Message: models.ChatMessage{
				Role:    choice.Message.Role,
				Content: choice.Message.Content,
			},
			FinishReason: string(choice.FinishReason),
		})
	}

	return &models.ChatCompletionResponse{
		ID:      resp.ID,
		Object:  resp.Object,
		Created: resp.Created,
		Model:   resp.Model,
		Choices: choices,
		Usage: models.Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}, nil
}

func (a *OpenAIAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	messages := make([]openai.ChatCompletionMessage, 0, len(req.Messages))
	for _, msg := range req.Messages {
		var content string
		switch v := msg.Content.(type) {
		case string:
			content = v
		default:
			content = fmt.Sprintf("%v", v)
		}
		
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: content,
		})
	}

	openaiReq := openai.ChatCompletionRequest{
		Model:       req.Model,
		Messages:    messages,
		Temperature: float32(getFloatValue(req.Temperature)),
		TopP:        float32(getFloatValue(req.TopP)),
		MaxTokens:   getIntValue(req.MaxTokens),
		Stream:      true,
	}

	stream, err := a.client.CreateChatCompletionStream(ctx, openaiReq)
	if err != nil {
		return nil, fmt.Errorf("openai stream error: %w", err)
	}

	return &OpenAIStreamReader{stream: stream}, nil
}

type OpenAIStreamReader struct {
	stream *openai.ChatCompletionStream
}

func (r *OpenAIStreamReader) Read(p []byte) (n int, err error) {
	// 这里需要将流式响应转换为 SSE 格式
	// 简化实现，实际应该处理 SSE 格式
	return 0, fmt.Errorf("stream reading not fully implemented")
}

func (r *OpenAIStreamReader) Close() error {
	return r.stream.Close()
}

// 辅助函数
func getFloatValue(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func getIntValue(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}


