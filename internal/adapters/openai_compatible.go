package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gotoailab/llmhub/internal/models"
)

// openAICompatibleAdapter 通用的 OpenAI 兼容适配器
type openAICompatibleAdapter struct {
	provider      Provider
	apiKey        string
	baseURL       string
	client        *http.Client
	endpoint      string // API 端点路径，默认为 /chat/completions
	authHeader    string // 认证头格式，默认为 "Bearer"
	supportsTools bool   // 是否支持工具调用
}

// NewOpenAICompatibleAdapter 创建通用的 OpenAI 兼容适配器
func NewOpenAICompatibleAdapter(provider Provider, apiKey, baseURL, endpoint, authHeader string) *openAICompatibleAdapter {
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	if endpoint == "" {
		endpoint = "/chat/completions"
	}
	if authHeader == "" {
		authHeader = "Bearer"
	}

	// 根据提供商判断是否支持工具调用
	supportsTools := isProviderSupportsTools(provider)

	return &openAICompatibleAdapter{
		provider:      provider,
		apiKey:        apiKey,
		baseURL:       baseURL,
		endpoint:      endpoint,
		authHeader:    authHeader,
		supportsTools: supportsTools,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (a *openAICompatibleAdapter) GetProvider() Provider {
	return a.provider
}

func (a *openAICompatibleAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	// 检查工具调用支持
	if (len(req.Tools) > 0 || len(req.Functions) > 0) && !a.supportsTools {
		return nil, fmt.Errorf("tool use not supported for provider %s", a.provider)
	}

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
	// 检查工具调用支持
	if (len(req.Tools) > 0 || len(req.Functions) > 0) && !a.supportsTools {
		return nil, fmt.Errorf("tool use not supported for provider %s", a.provider)
	}

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

// 判断提供商是否支持工具调用
func isProviderSupportsTools(provider Provider) bool {
	supportedProviders := map[Provider]bool{
		"openrouter":   true,
		"groq":         true,
		"together":     true,
		"deepseek":     true,
		"siliconflow":  true,
		"moonshot":     true,
		"stepfun":      true,
		"mistral":      true,
		"cohere":       true,
		"qwen":         false, // 通义千问暂不支持标准工具调用
		"baichuan":     false,
		"chatglm":      false,
		"ernie":        false,
		"spark":        false,
		"hunyuan":      false,
		"360":          false,
		"minimax":      false,
		"yi":           false,
		"doubao":       false,
		"novita":       true,
		"xai":          true,
		"ollama":       false, // Ollama 需要特殊处理
		"coze":         false,
	}

	supported, exists := supportedProviders[provider]
	if !exists {
		// 对于未知提供商，默认不支持
		return false
	}
	return supported
}
