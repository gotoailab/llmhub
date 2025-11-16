package adapters

import (
	"context"
	"io"

	"github.com/gotoailab/llmhub/internal/models"
)

// XaiAdapter xAI适配器
type XaiAdapter struct {
	*openAICompatibleAdapter
}

// NewXaiAdapter 创建xAI适配器
func NewXaiAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.x.ai/v1"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("xai"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &XaiAdapter{adapter}, nil
}

func (a *XaiAdapter) GetProvider() Provider {
	return Provider("xai")
}

func (a *XaiAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *XaiAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
