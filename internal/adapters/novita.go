package adapters

import (
	"github.com/aihub/internal/models"
	"context"
	"io"

)

// NovitaAdapter novita.ai适配器
type NovitaAdapter struct {
	*openAICompatibleAdapter
}

// NewNovitaAdapter 创建novita.ai适配器
func NewNovitaAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.novita.ai/v3"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("novita"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &NovitaAdapter{adapter}, nil
}

func (a *NovitaAdapter) GetProvider() Provider {
	return Provider("novita")
}

func (a *NovitaAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *NovitaAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
