package adapters

import (
	"context"
	"io"

	"github.com/gotoailab/llmhub/internal/models"
)

// YiAdapter 零一万物适配器
type YiAdapter struct {
	*openAICompatibleAdapter
}

// NewYiAdapter 创建零一万物适配器
func NewYiAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.lingyiwanwu.com/v1"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("yi"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &YiAdapter{adapter}, nil
}

func (a *YiAdapter) GetProvider() Provider {
	return Provider("yi")
}

func (a *YiAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *YiAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
