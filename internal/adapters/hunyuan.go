package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// HunyuanAdapter 腾讯混元适配器
type HunyuanAdapter struct {
	*openAICompatibleAdapter
}

// NewHunyuanAdapter 创建腾讯混元适配器
func NewHunyuanAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://hunyuan.tencentcloudapi.com"
	}
	adapter := NewOpenAICompatibleAdapter("hunyuan", apiKey, baseURL, "/chat/completions", "Bearer")
	return &HunyuanAdapter{adapter}, nil
}

func (a *HunyuanAdapter) GetProvider() string {
	return "hunyuan"
}

func (a *HunyuanAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *HunyuanAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
