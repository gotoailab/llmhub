package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// MoonshotAdapter Moonshot AI适配器
type MoonshotAdapter struct {
	*openAICompatibleAdapter
}

// NewMoonshotAdapter 创建Moonshot AI适配器
func NewMoonshotAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.moonshot.cn/v1"
	}
	adapter := NewOpenAICompatibleAdapter("moonshot", apiKey, baseURL, "/chat/completions", "Bearer")
	return &MoonshotAdapter{adapter}, nil
}

func (a *MoonshotAdapter) GetProvider() string {
	return "moonshot"
}

func (a *MoonshotAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *MoonshotAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
