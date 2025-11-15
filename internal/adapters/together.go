package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// TogetherAdapter together.ai适配器
type TogetherAdapter struct {
	*openAICompatibleAdapter
}

// NewTogetherAdapter 创建together.ai适配器
func NewTogetherAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.together.xyz/v1"
	}
	adapter := NewOpenAICompatibleAdapter("together", apiKey, baseURL, "/chat/completions", "Bearer")
	return &TogetherAdapter{adapter}, nil
}

func (a *TogetherAdapter) GetProvider() string {
	return "together"
}

func (a *TogetherAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *TogetherAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
