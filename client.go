package aihub

import (
	"context"
	"fmt"
	"io"

	"github.com/aihub/internal/adapters"
)

// Provider 支持的模型提供商
type Provider string

const (
	ProviderOpenAI      Provider = "openai"
	ProviderClaude      Provider = "claude"
	ProviderDeepSeek    Provider = "deepseek"
	ProviderQwen        Provider = "qwen"
	ProviderSiliconFlow Provider = "siliconflow"
	ProviderGemini      Provider = "gemini"
	ProviderMistral     Provider = "mistral"
	ProviderDoubao      Provider = "doubao"
	ProviderErnie       Provider = "ernie"
	ProviderSpark       Provider = "spark"
	ProviderChatGLM     Provider = "chatglm"
	Provider360         Provider = "360"
	ProviderHunyuan     Provider = "hunyuan"
	ProviderMoonshot    Provider = "moonshot"
	ProviderBaichuan    Provider = "baichuan"
	ProviderMiniMax     Provider = "minimax"
	ProviderGroq        Provider = "groq"
	ProviderOllama      Provider = "ollama"
	ProviderYi          Provider = "yi"
	ProviderStepFun     Provider = "stepfun"
	ProviderCoze        Provider = "coze"
	ProviderCohere      Provider = "cohere"
	ProviderCloudflare  Provider = "cloudflare"
	ProviderDeepL       Provider = "deepl"
	ProviderTogether    Provider = "together"
	ProviderNovita      Provider = "novita"
	ProviderXAI         Provider = "xai"
)

// Client 客户端接口，提供 OpenAI 兼容的方法
type Client struct {
	adapter *adapterWrapper
	model   string
}

// ClientConfig 客户端配置
type ClientConfig struct {
	// APIKey 模型提供商的 API Key
	APIKey string

	// Provider 模型提供商
	Provider Provider

	// BaseURL 可选的 API 基础 URL，如果不提供则使用默认值
	BaseURL string

	// Model 模型名称（可选，可以在调用时指定）
	Model string
}

// NewClient 创建新的客户端
// 使用示例：
//
//	client := aihub.NewClient(aihub.ClientConfig{
//	    APIKey: "sk-your-api-key",
//	    Provider: aihub.ProviderOpenAI,
//	    Model: "gpt-3.5-turbo",
//	})
func NewClient(config ClientConfig) (*Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("api key is required")
	}

	if config.Provider == "" {
		return nil, fmt.Errorf("provider is required")
	}

	// 创建适配器
	adapter, err := adapters.CreateAdapter(string(config.Provider), config.APIKey, config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create adapter: %w", err)
	}

	return &Client{
		adapter: &adapterWrapper{adapter: adapter},
		model:   config.Model,
	}, nil
}

// ChatCompletions 创建聊天完成请求（非流式）
func (c *Client) ChatCompletions(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionResponse, error) {
	// 如果请求中没有指定模型，使用客户端默认模型
	if req.Model == "" {
		if c.model == "" {
			return nil, fmt.Errorf("model is required")
		}
		req.Model = c.model
	}

	// 转换为内部请求格式
	internalReq := c.toInternalRequest(req)

	// 调用适配器包装器
	internalResp, err := c.adapter.ChatCompletion(ctx, internalReq)
	if err != nil {
		return nil, err
	}

	// 转换为公共响应格式
	return c.toPublicResponse(internalResp), nil
}

// ChatCompletionsStream 创建流式聊天完成请求
func (c *Client) ChatCompletionsStream(ctx context.Context, req ChatCompletionRequest) (io.ReadCloser, error) {
	// 如果请求中没有指定模型，使用客户端默认模型
	if req.Model == "" {
		if c.model == "" {
			return nil, fmt.Errorf("model is required")
		}
		req.Model = c.model
	}

	// 设置流式标志
	req.Stream = true

	// 转换为内部请求格式
	internalReq := c.toInternalRequest(req)

	// 调用适配器包装器
	return c.adapter.ChatCompletionStream(ctx, internalReq)
}

// GetProvider 获取当前客户端使用的提供商
func (c *Client) GetProvider() string {
	return c.adapter.GetProvider()
}

// toInternalRequest 转换为内部请求格式
func (c *Client) toInternalRequest(req ChatCompletionRequest) *internalChatCompletionRequest {
	return &internalChatCompletionRequest{
		Model:            req.Model,
		Messages:         c.toInternalMessages(req.Messages),
		Temperature:      req.Temperature,
		TopP:             req.TopP,
		MaxTokens:        req.MaxTokens,
		Stream:           req.Stream,
		PresencePenalty:  req.PresencePenalty,
		FrequencyPenalty: req.FrequencyPenalty,
		Stop:             req.Stop,
		User:             req.User,
		Functions:        c.toInternalFunctions(req.Functions),
		FunctionCall:     req.FunctionCall,
		LogitBias:        req.LogitBias,
		LogProbs:         req.LogProbs,
		TopLogProbs:      req.TopLogProbs,
		ResponseFormat:   c.toInternalResponseFormat(req.ResponseFormat),
		Seed:             req.Seed,
		Tools:            c.toInternalTools(req.Tools),
		ToolChoice:       req.ToolChoice,
	}
}

// toPublicResponse 转换为公共响应格式
func (c *Client) toPublicResponse(resp *internalChatCompletionResponse) *ChatCompletionResponse {
	choices := make([]ChatCompletionChoice, 0, len(resp.Choices))
	for _, choice := range resp.Choices {
		choices = append(choices, ChatCompletionChoice{
			Index:        choice.Index,
			Message:      c.toPublicMessage(choice.Message),
			FinishReason: choice.FinishReason,
			Delta:        c.toPublicMessagePtr(choice.Delta),
		})
	}

	return &ChatCompletionResponse{
		ID:                resp.ID,
		Object:            resp.Object,
		Created:           resp.Created,
		Model:             resp.Model,
		Choices:           choices,
		Usage:             Usage(resp.Usage),
		SystemFingerprint: resp.SystemFingerprint,
	}
}

// 类型转换辅助函数
func (c *Client) toInternalMessages(msgs []ChatMessage) []internalChatMessage {
	result := make([]internalChatMessage, 0, len(msgs))
	for _, msg := range msgs {
		result = append(result, internalChatMessage{
			Role:         msg.Role,
			Content:      msg.Content,
			Name:         msg.Name,
			FunctionCall: c.toInternalFunctionCall(msg.FunctionCall),
			ToolCalls:    c.toInternalToolCalls(msg.ToolCalls),
			ToolCallID:   msg.ToolCallID,
		})
	}
	return result
}

func (c *Client) toPublicMessage(msg internalChatMessage) ChatMessage {
	return ChatMessage{
		Role:         msg.Role,
		Content:      msg.Content,
		Name:         msg.Name,
		FunctionCall: c.toPublicFunctionCall(msg.FunctionCall),
		ToolCalls:    c.toPublicToolCalls(msg.ToolCalls),
		ToolCallID:   msg.ToolCallID,
	}
}

func (c *Client) toPublicMessagePtr(msg *internalChatMessage) *ChatMessage {
	if msg == nil {
		return nil
	}
	m := c.toPublicMessage(*msg)
	return &m
}

func (c *Client) toInternalFunctions(funcs []FunctionDefinition) []internalFunctionDefinition {
	result := make([]internalFunctionDefinition, 0, len(funcs))
	for _, f := range funcs {
		result = append(result, internalFunctionDefinition{
			Name:        f.Name,
			Description: f.Description,
			Parameters:  f.Parameters,
		})
	}
	return result
}

func (c *Client) toPublicFunctionCall(fc *internalFunctionCall) *FunctionCall {
	if fc == nil {
		return nil
	}
	return &FunctionCall{
		Name:      fc.Name,
		Arguments: fc.Arguments,
	}
}

func (c *Client) toInternalFunctionCall(fc *FunctionCall) *internalFunctionCall {
	if fc == nil {
		return nil
	}
	return &internalFunctionCall{
		Name:      fc.Name,
		Arguments: fc.Arguments,
	}
}

func (c *Client) toInternalToolCalls(tcs []ToolCall) []internalToolCall {
	result := make([]internalToolCall, 0, len(tcs))
	for _, tc := range tcs {
		result = append(result, internalToolCall{
			ID:       tc.ID,
			Type:     tc.Type,
			Function: *c.toInternalFunctionCall(&tc.Function),
		})
	}
	return result
}

func (c *Client) toPublicToolCalls(tcs []internalToolCall) []ToolCall {
	result := make([]ToolCall, 0, len(tcs))
	for _, tc := range tcs {
		result = append(result, ToolCall{
			ID:       tc.ID,
			Type:     tc.Type,
			Function: *c.toPublicFunctionCall(&tc.Function),
		})
	}
	return result
}

func (c *Client) toInternalTools(tools []Tool) []internalTool {
	result := make([]internalTool, 0, len(tools))
	for _, tool := range tools {
		result = append(result, internalTool{
			Type:     tool.Type,
			Function: c.toInternalFunctions([]FunctionDefinition{tool.Function})[0],
		})
	}
	return result
}

func (c *Client) toInternalResponseFormat(rf *ResponseFormat) *internalResponseFormat {
	if rf == nil {
		return nil
	}
	return &internalResponseFormat{
		Type: rf.Type,
	}
}
