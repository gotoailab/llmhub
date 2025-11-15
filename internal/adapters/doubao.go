package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// DoubaoAdapter 豆包（字节跳动）适配器
type DoubaoAdapter struct {
	*openAICompatibleAdapter
}

// NewDoubaoAdapter 创建豆包适配器
func NewDoubaoAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://ark.cn-beijing.volces.com/api/v3"
	}
	adapter := NewOpenAICompatibleAdapter("doubao", apiKey, baseURL, "/chat/completions", "Bearer")
	return &DoubaoAdapter{adapter}, nil
}

func (a *DoubaoAdapter) GetProvider() string {
	return "doubao"
}

func (a *DoubaoAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *DoubaoAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}

