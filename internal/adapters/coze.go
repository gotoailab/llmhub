package adapters

import (
	"github.com/aihub/internal/models"
	"context"
	"io"

)

// CozeAdapter Coze适配器
type CozeAdapter struct {
	*openAICompatibleAdapter
}

// NewCozeAdapter 创建Coze适配器
func NewCozeAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.coze.cn/v1"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("coze"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &CozeAdapter{adapter}, nil
}

func (a *CozeAdapter) GetProvider() Provider {
	return Provider("coze")
}

func (a *CozeAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *CozeAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
