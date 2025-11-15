package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// MinimaxAdapter MINIMAX适配器
type MinimaxAdapter struct {
	*openAICompatibleAdapter
}

// NewMinimaxAdapter 创建MINIMAX适配器
func NewMinimaxAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.minimax.chat/v1"
	}
	adapter := NewOpenAICompatibleAdapter("minimax", apiKey, baseURL, "/chat/completions", "Bearer")
	return &MinimaxAdapter{adapter}, nil
}

func (a *MinimaxAdapter) GetProvider() string {
	return "minimax"
}

func (a *MinimaxAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *MinimaxAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
