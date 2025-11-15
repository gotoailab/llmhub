package adapters

import (
	"github.com/aihub/internal/models"
	"context"
	"io"

)

// BaichuanAdapter 百川大模型适配器
type BaichuanAdapter struct {
	*openAICompatibleAdapter
}

// NewBaichuanAdapter 创建百川大模型适配器
func NewBaichuanAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.baichuan-ai.com/v1"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("baichuan"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &BaichuanAdapter{adapter}, nil
}

func (a *BaichuanAdapter) GetProvider() Provider {
	return Provider("baichuan")
}

func (a *BaichuanAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *BaichuanAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
