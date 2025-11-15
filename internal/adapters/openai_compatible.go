package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aihub/internal/models"
)

// openAICompatibleAdapter 通用的 OpenAI 兼容适配器
type openAICompatibleAdapter struct {
	provider   string
	apiKey     string
	baseURL    string
	client     *http.Client
	endpoint   string // API 端点路径，默认为 /chat/completions
	authHeader string // 认证头格式，默认为 "Bearer"
}

// NewOpenAICompatibleAdapter 创建通用的 OpenAI 兼容适配器
func NewOpenAICompatibleAdapter(provider, apiKey, baseURL, endpoint, authHeader string) *openAICompatibleAdapter {
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	if endpoint == "" {
		endpoint = "/chat/completions"
	}
	if authHeader == "" {
		authHeader = "Bearer"
	}

	return &openAICompatibleAdapter{
		provider:   provider,
		apiKey:     apiKey,
		baseURL:    baseURL,
		endpoint:   endpoint,
		authHeader: authHeader,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (a *openAICompatibleAdapter) GetProvider() string {
	return a.provider
}

func (a *openAICompatibleAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	openaiReq := convertToOpenAIFormatGeneric(req)

	reqBody, err := json.Marshal(openaiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+a.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if a.apiKey != "" {
		if a.authHeader == "Bearer" {
			httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)
		} else {
			httpReq.Header.Set("Authorization", a.apiKey)
		}
	}

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("%s api error: %w", a.provider, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s api error: status %d, body: %s", a.provider, resp.StatusCode, string(body))
	}

	var openaiResp models.ChatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &openaiResp, nil
}

func (a *openAICompatibleAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	openaiReq := convertToOpenAIFormatGeneric(req)
	openaiReq["stream"] = true

	reqBody, err := json.Marshal(openaiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+a.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if a.apiKey != "" {
		if a.authHeader == "Bearer" {
			httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)
		} else {
			httpReq.Header.Set("Authorization", a.apiKey)
		}
	}

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("%s stream error: %w", a.provider, err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("%s stream error: status %d", a.provider, resp.StatusCode)
	}

	return resp.Body, nil
}

