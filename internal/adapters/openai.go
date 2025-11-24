package adapters

import (
	"context"
	"fmt"
	"io"

	"github.com/gotoailab/llmhub/internal/models"

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
		openaiMsg := openai.ChatCompletionMessage{
			Role: msg.Role,
			Name: msg.Name,
		}

		// 处理消息内容
		switch v := msg.Content.(type) {
		case string:
			openaiMsg.Content = v
		case []interface{}:
			// 处理多模态内容
			openaiMsg.Content = fmt.Sprintf("%v", v)
		default:
			openaiMsg.Content = fmt.Sprintf("%v", v)
		}

		// 处理工具调用
		if len(msg.ToolCalls) > 0 {
			toolCalls := make([]openai.ToolCall, 0, len(msg.ToolCalls))
			for _, tc := range msg.ToolCalls {
				toolCalls = append(toolCalls, openai.ToolCall{
					ID:   tc.ID,
					Type: openai.ToolType(tc.Type),
					Function: openai.FunctionCall{
						Name:      tc.Function.Name,
						Arguments: tc.Function.Arguments,
					},
				})
			}
			openaiMsg.ToolCalls = toolCalls
		}

		// 处理函数调用（旧格式兼容）
		if msg.FunctionCall != nil {
			openaiMsg.FunctionCall = &openai.FunctionCall{
				Name:      msg.FunctionCall.Name,
				Arguments: msg.FunctionCall.Arguments,
			}
		}

		// 处理工具调用 ID
		if msg.ToolCallID != "" {
			openaiMsg.ToolCallID = msg.ToolCallID
		}

		messages = append(messages, openaiMsg)
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

	// 添加工具支持
	if len(req.Tools) > 0 {
		tools := make([]openai.Tool, 0, len(req.Tools))
		for _, tool := range req.Tools {
			tools = append(tools, openai.Tool{
				Type: openai.ToolType(tool.Type),
				Function: &openai.FunctionDefinition{
					Name:        tool.Function.Name,
					Description: tool.Function.Description,
					Parameters:  tool.Function.Parameters,
				},
			})
		}
		openaiReq.Tools = tools

		// 处理工具选择策略
		if req.ToolChoice != nil {
			openaiReq.ToolChoice = req.ToolChoice
		}
	}

	// 添加函数支持（旧格式兼容）
	if len(req.Functions) > 0 {
		functions := make([]openai.FunctionDefinition, 0, len(req.Functions))
		for _, fn := range req.Functions {
			functions = append(functions, openai.FunctionDefinition{
				Name:        fn.Name,
				Description: fn.Description,
				Parameters:  fn.Parameters,
			})
		}
		openaiReq.Functions = functions

		if req.FunctionCall != nil {
			openaiReq.FunctionCall = req.FunctionCall
		}
	}

	// 调用 OpenAI API
	resp, err := a.client.CreateChatCompletion(ctx, openaiReq)
	if err != nil {
		return nil, fmt.Errorf("openai api error: %w", err)
	}

	// 转换响应格式
	choices := make([]models.ChatCompletionChoice, 0, len(resp.Choices))
	for _, choice := range resp.Choices {
		message := models.ChatMessage{
			Role:    choice.Message.Role,
			Content: choice.Message.Content,
			Name:    choice.Message.Name,
		}

		// 处理工具调用响应
		if len(choice.Message.ToolCalls) > 0 {
			toolCalls := make([]models.ToolCall, 0, len(choice.Message.ToolCalls))
			for _, tc := range choice.Message.ToolCalls {
				toolCalls = append(toolCalls, models.ToolCall{
					ID:   tc.ID,
					Type: string(tc.Type),
					Function: models.FunctionCall{
						Name:      tc.Function.Name,
						Arguments: tc.Function.Arguments,
					},
				})
			}
			message.ToolCalls = toolCalls
		}

		// 处理函数调用响应（旧格式兼容）
		if choice.Message.FunctionCall != nil {
			message.FunctionCall = &models.FunctionCall{
				Name:      choice.Message.FunctionCall.Name,
				Arguments: choice.Message.FunctionCall.Arguments,
			}
		}

		choices = append(choices, models.ChatCompletionChoice{
			Index:        choice.Index,
			Message:      message,
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
		openaiMsg := openai.ChatCompletionMessage{
			Role: msg.Role,
			Name: msg.Name,
		}

		switch v := msg.Content.(type) {
		case string:
			openaiMsg.Content = v
		default:
			openaiMsg.Content = fmt.Sprintf("%v", v)
		}

		// 处理工具调用和函数调用（与上面类似的逻辑）
		if len(msg.ToolCalls) > 0 {
			toolCalls := make([]openai.ToolCall, 0, len(msg.ToolCalls))
			for _, tc := range msg.ToolCalls {
				toolCalls = append(toolCalls, openai.ToolCall{
					ID:   tc.ID,
					Type: openai.ToolType(tc.Type),
					Function: openai.FunctionCall{
						Name:      tc.Function.Name,
						Arguments: tc.Function.Arguments,
					},
				})
			}
			openaiMsg.ToolCalls = toolCalls
		}

		if msg.FunctionCall != nil {
			openaiMsg.FunctionCall = &openai.FunctionCall{
				Name:      msg.FunctionCall.Name,
				Arguments: msg.FunctionCall.Arguments,
			}
		}

		if msg.ToolCallID != "" {
			openaiMsg.ToolCallID = msg.ToolCallID
		}

		messages = append(messages, openaiMsg)
	}

	openaiReq := openai.ChatCompletionRequest{
		Model:       req.Model,
		Messages:    messages,
		Temperature: float32(getFloatValue(req.Temperature)),
		TopP:        float32(getFloatValue(req.TopP)),
		MaxTokens:   getIntValue(req.MaxTokens),
		Stream:      true,
	}

	// 添加工具和函数支持（与上面相同的逻辑）
	if len(req.Tools) > 0 {
		tools := make([]openai.Tool, 0, len(req.Tools))
		for _, tool := range req.Tools {
			tools = append(tools, openai.Tool{
				Type: openai.ToolType(tool.Type),
				Function: &openai.FunctionDefinition{
					Name:        tool.Function.Name,
					Description: tool.Function.Description,
					Parameters:  tool.Function.Parameters,
				},
			})
		}
		openaiReq.Tools = tools
		if req.ToolChoice != nil {
			openaiReq.ToolChoice = req.ToolChoice
		}
	}

	if len(req.Functions) > 0 {
		functions := make([]openai.FunctionDefinition, 0, len(req.Functions))
		for _, fn := range req.Functions {
			functions = append(functions, openai.FunctionDefinition{
				Name:        fn.Name,
				Description: fn.Description,
				Parameters:  fn.Parameters,
			})
		}
		openaiReq.Functions = functions
		if req.FunctionCall != nil {
			openaiReq.FunctionCall = req.FunctionCall
		}
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
