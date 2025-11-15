package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// StepfunAdapter 阶跃星辰适配器
type StepfunAdapter struct {
	*openAICompatibleAdapter
}

// NewStepfunAdapter 创建阶跃星辰适配器
func NewStepfunAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.stepfun.com/v1"
	}
	adapter := NewOpenAICompatibleAdapter("stepfun", apiKey, baseURL, "/chat/completions", "Bearer")
	return &StepfunAdapter{adapter}, nil
}

func (a *StepfunAdapter) GetProvider() string {
	return "stepfun"
}

func (a *StepfunAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *StepfunAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
