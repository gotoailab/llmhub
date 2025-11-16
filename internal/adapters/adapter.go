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
