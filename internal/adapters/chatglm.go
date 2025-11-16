package adapters

import (
	"context"
	"io"

	"github.com/gotoailab/llmhub/internal/models"
)

// ChatglmAdapter ChatGLM（智谱）适配器
type ChatglmAdapter struct {
	*openAICompatibleAdapter
}

// NewChatglmAdapter 创建ChatGLM（智谱）适配器
func NewChatglmAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://open.bigmodel.cn/api/paas/v4"
	}
	adapter := NewOpenAICompatibleAdapter(Provider("chatglm"), apiKey, baseURL, "/chat/completions", "Bearer")
	return &ChatglmAdapter{adapter}, nil
}

func (a *ChatglmAdapter) GetProvider() Provider {
	return Provider("chatglm")
}

func (a *ChatglmAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *ChatglmAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
