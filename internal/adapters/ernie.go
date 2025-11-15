package adapters

import (
	"github.com/aihub/internal/models"
	"context"
	"io"

)

// ErnieAdapter 文心一言（百度）适配器
type ErnieAdapter struct {
	*openAICompatibleAdapter
}

// NewErnieAdapter 创建文心一言（百度）适配器
func NewErnieAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("ernie"), apiKey, baseURL, "/completions", "Bearer")
	return &ErnieAdapter{adapter}, nil
}

func (a *ErnieAdapter) GetProvider() Provider {
	return Provider("ernie")
}

func (a *ErnieAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *ErnieAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
