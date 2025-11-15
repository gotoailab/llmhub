package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// OllamaAdapter Ollama适配器
type OllamaAdapter struct {
	*openAICompatibleAdapter
}

// NewOllamaAdapter 创建Ollama适配器
func NewOllamaAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "http://localhost:11434"
	}
	adapter := NewOpenAICompatibleAdapter("ollama", apiKey, baseURL, "/api/chat", "")
	return &OllamaAdapter{adapter}, nil
}

func (a *OllamaAdapter) GetProvider() string {
	return "ollama"
}

func (a *OllamaAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *OllamaAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
