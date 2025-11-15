package adapters

import (
	"context"
	"io"

	"github.com/aihub/internal/models"
)

// SparkAdapter 讯飞星火适配器
type SparkAdapter struct {
	*openAICompatibleAdapter
}

// NewSparkAdapter 创建讯飞星火适配器
func NewSparkAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://spark-api.xf-yun.com/v1"
	}
	adapter := NewOpenAICompatibleAdapter("spark", apiKey, baseURL, "/chat/completions", "Bearer")
	return &SparkAdapter{adapter}, nil
}

func (a *SparkAdapter) GetProvider() string {
	return "spark"
}

func (a *SparkAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return a.openAICompatibleAdapter.ChatCompletion(ctx, req)
}

func (a *SparkAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	return a.openAICompatibleAdapter.ChatCompletionStream(ctx, req)
}
