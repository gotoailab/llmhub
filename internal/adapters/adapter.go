package adapters

import (
	"context"
	"fmt"
	"io"

	"github.com/gotoailab/llmhub/internal/models"
)

// Provider 提供商类型（与 Provider 兼容）
type Provider string

// Adapter 定义所有模型适配器必须实现的接口
type Adapter interface {
	// ChatCompletion 处理聊天完成请求
	ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error)

	// ChatCompletionStream 处理流式聊天完成请求
	ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error)

	// GetProvider 返回适配器对应的提供商
	GetProvider() Provider
}

// AdapterFactory 创建适配器的工厂函数
type AdapterFactory func(apiKey, baseURL string) (Adapter, error)

// Registry 适配器注册表，使用 Provider 作为 key
var Registry = make(map[Provider]AdapterFactory)

// Register 注册适配器工厂
func Register(provider Provider, factory AdapterFactory) {
	Registry[provider] = factory
}

// CreateAdapter 根据提供商创建适配器
func CreateAdapter(provider Provider, apiKey, baseURL string) (Adapter, error) {
	factory, exists := Registry[provider]
	if !exists {
		return nil, fmt.Errorf("provider %s is not registered", provider)
	}
	return factory(apiKey, baseURL)
}

// convertToOpenAIFormatGeneric 通用的 OpenAI 格式转换函数
func convertToOpenAIFormatGeneric(req *models.ChatCompletionRequest) map[string]interface{} {
	result := map[string]interface{}{
		"model":    req.Model,
		"messages": req.Messages,
	}

	if req.Temperature != nil {
		result["temperature"] = *req.Temperature
	}
	if req.TopP != nil {
		result["top_p"] = *req.TopP
	}
	if req.MaxTokens != nil {
		result["max_tokens"] = *req.MaxTokens
	}
	if req.PresencePenalty != nil {
		result["presence_penalty"] = *req.PresencePenalty
	}
	if req.FrequencyPenalty != nil {
		result["frequency_penalty"] = *req.FrequencyPenalty
	}
	if len(req.Stop) > 0 {
		result["stop"] = req.Stop
	}
	if req.User != "" {
		result["user"] = req.User
	}

	// 添加工具支持
	if len(req.Tools) > 0 {
		result["tools"] = req.Tools
		if req.ToolChoice != nil {
			result["tool_choice"] = req.ToolChoice
		}
	}

	// 添加函数支持（旧格式）
	if len(req.Functions) > 0 {
		result["functions"] = req.Functions
		if req.FunctionCall != nil {
			result["function_call"] = req.FunctionCall
		}
	}

	// 添加其他参数
	if req.LogitBias != nil {
		result["logit_bias"] = req.LogitBias
	}
	if req.LogProbs {
		result["logprobs"] = req.LogProbs
		if req.TopLogProbs != nil {
			result["top_logprobs"] = *req.TopLogProbs
		}
	}
	if req.ResponseFormat != nil {
		result["response_format"] = req.ResponseFormat
	}
	if req.Seed != nil {
		result["seed"] = *req.Seed
	}

	return result
}
