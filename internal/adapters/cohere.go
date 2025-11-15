package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// CohereAdapter Cohere适配器
type CohereAdapter struct {
	*openAICompatibleAdapter
}

// NewCohereAdapter 创建Cohere适配器
func NewCohereAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.cohere.ai/v1"
	}
	adapter := NewOpenAICompatibleAdapter("cohere", apiKey, baseURL, "/chat/completions", "Bearer")
	return &CohereAdapter{adapter}, nil
}

func (a *CohereAdapter) GetProvider() string {
	return "cohere"
}

func (a *CohereAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *CohereAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
