package adapters

import (
	"context"
	"io"

	"github.com/gotoailab/llmhub/internal/models"
)

// GroqAdapter Groq适配器
type GroqAdapter struct {
	*openAICompatibleAdapter
}

// NewGroqAdapter 创建Groq适配器
func NewGroqAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.groq.com/openai/v1"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("groq"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &GroqAdapter{adapter}, nil
}

func (a *GroqAdapter) GetProvider() Provider {
	return Provider("groq")
}

func (a *GroqAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *GroqAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
