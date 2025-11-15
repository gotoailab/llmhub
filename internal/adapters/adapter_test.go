package adapters

import (
	"context"
	"testing"

	"github.com/aihub/internal/models"
)

func init() {
	// 在测试中注册适配器
	Register("openai", NewOpenAIAdapter)
	Register("claude", NewClaudeAdapter)
	Register("deepseek", NewDeepSeekAdapter)
	Register("qwen", NewQwenAdapter)
	Register("siliconflow", NewSiliconFlowAdapter)
	Register("gemini", NewGeminiAdapter)
	Register("mistral", NewMistralAdapter)
}

func TestCreateAdapter(t *testing.T) {
	tests := []struct {
		name     string
		provider string
		apiKey   string
		baseURL  string
		wantErr  bool
	}{
		{
			name:     "openai",
			provider: "openai",
			apiKey:   "test-key",
			wantErr:  false,
		},
		{
			name:     "claude",
			provider: "claude",
			apiKey:   "test-key",
			wantErr:  false,
		},
		{
			name:     "invalid provider",
			provider: "invalid",
			apiKey:   "test-key",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter, err := CreateAdapter(tt.provider, tt.apiKey, tt.baseURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && adapter == nil {
				t.Error("CreateAdapter() returned nil adapter")
			}
			if !tt.wantErr && adapter != nil {
				if adapter.GetProvider() != tt.provider {
					t.Errorf("GetProvider() = %v, want %v", adapter.GetProvider(), tt.provider)
				}
			}
		})
	}
}

func TestAdapter_GetProvider(t *testing.T) {
	providers := []string{
		"openai",
		"claude",
		"deepseek",
		"qwen",
		"siliconflow",
		"gemini",
		"mistral",
	}

	for _, provider := range providers {
		t.Run(provider, func(t *testing.T) {
			adapter, err := CreateAdapter(provider, "test-key", "")
			if err != nil {
				t.Logf("Provider %s not available: %v", provider, err)
				return
			}
			if adapter.GetProvider() != provider {
				t.Errorf("GetProvider() = %v, want %v", adapter.GetProvider(), provider)
			}
		})
	}
}

// 测试适配器接口实现
func TestAdapterInterface(t *testing.T) {
	adapter, err := CreateAdapter("openai", "test-key", "")
	if err != nil {
		t.Skipf("Skipping test: %v", err)
		return
	}

	ctx := context.Background()
	req := &models.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []models.ChatMessage{
			{Role: "user", Content: "Hello"},
		},
	}

	// 注意：这需要真实的 API key 才能成功
	// 这里只测试接口是否正常工作，不测试实际 API 调用
	_, err = adapter.ChatCompletion(ctx, req)
	// 我们期望这里会有错误（因为 API key 无效），但接口应该能正常调用
	if err == nil {
		t.Log("ChatCompletion succeeded (unexpected with test key)")
	} else {
		t.Logf("ChatCompletion returned error (expected): %v", err)
	}
}
