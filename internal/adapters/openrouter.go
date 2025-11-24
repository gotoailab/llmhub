package adapters

import (
	"context"
	"io"

	"github.com/gotoailab/llmhub/internal/models"
)

// OpenRouterAdapter OpenRouter适配器
type OpenRouterAdapter struct {
	*openAICompatibleAdapter
}

// NewOpenRouterAdapter 创建OpenRouter适配器
func NewOpenRouterAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://openrouter.ai/api/v1"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("openrouter"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &OpenRouterAdapter{adapter}, nil
}

func (a *OpenRouterAdapter) GetProvider() Provider {
	return Provider("openrouter")
}

func (a *OpenRouterAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *OpenRouterAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}