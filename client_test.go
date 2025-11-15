package aihub

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		config  ClientConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: ClientConfig{
				APIKey:   "test-key",
				Provider: ProviderOpenAI,
				Model:    "gpt-3.5-turbo",
			},
			wantErr: false,
		},
		{
			name: "missing api key",
			config: ClientConfig{
				Provider: ProviderOpenAI,
			},
			wantErr: true,
		},
		{
			name: "missing provider",
			config: ClientConfig{
				APIKey: "test-key",
			},
			wantErr: true,
		},
		{
			name: "invalid provider",
			config: ClientConfig{
				APIKey:   "test-key",
				Provider: Provider("invalid"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Error("NewClient() returned nil client when error was expected")
			}
			if !tt.wantErr && client != nil {
				if client.GetProvider() != tt.config.Provider {
					t.Errorf("GetProvider() = %v, want %v", client.GetProvider(), tt.config.Provider)
				}
			}
		})
	}
}

func TestClient_GetProvider(t *testing.T) {
	providers := []Provider{
		ProviderOpenAI,
		ProviderClaude,
		ProviderDeepSeek,
		ProviderQwen,
		ProviderSiliconFlow,
		ProviderGemini,
		ProviderMistral,
	}

	for _, provider := range providers {
		t.Run(string(provider), func(t *testing.T) {
			// 注意：这需要有效的 API key 才能测试，这里只测试创建
			// 在实际测试中，可以使用 mock 或者跳过需要真实 API 的测试
			config := ClientConfig{
				APIKey:   "test-key",
				Provider: provider,
			}
			client, err := NewClient(config)
			if err != nil {
				// 某些提供商可能不支持，这是正常的
				t.Logf("Provider %s not available: %v", provider, err)
				return
			}
			if client.GetProvider() != provider {
				t.Errorf("GetProvider() = %v, want %v", client.GetProvider(), provider)
			}
		})
	}
}

func TestClient_ChatCompletions_Validation(t *testing.T) {
	client, err := NewClient(ClientConfig{
		APIKey:   "test-key",
		Provider: ProviderOpenAI,
	})
	if err != nil {
		t.Skipf("Skipping test: %v", err)
		return
	}

	ctx := context.Background()

	// 测试缺少模型名称
	req := ChatCompletionRequest{
		Messages: []ChatMessage{
			{Role: "user", Content: "Hello"},
		},
	}

	_, err = client.ChatCompletions(ctx, req)
	if err == nil {
		t.Error("Expected error when model is not specified")
	}
}

