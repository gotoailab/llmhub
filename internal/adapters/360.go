package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// Three60Adapter 360智脑适配器
type Three60Adapter struct {
	*openAICompatibleAdapter
}

// New360Adapter 创建360智脑适配器
func New360Adapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.360.cn/v1"
	}
	adapter := NewOpenAICompatibleAdapter("360", apiKey, baseURL, "/chat/completions", "Bearer")
	return &Three60Adapter{adapter}, nil
}

func (a *Three60Adapter) GetProvider() string {
	return "360"
}

func (a *Three60Adapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *Three60Adapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
