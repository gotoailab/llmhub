package adapters

import (
	"github.com/aihub/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

)

type SiliconFlowAdapter struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewSiliconFlowAdapter 创建硅基流动适配器（导出以供注册）
func NewSiliconFlowAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.siliconflow.cn/v1"
	}
	
	return &SiliconFlowAdapter{
		apiKey:  apiKey,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}, nil
}

func (a *SiliconFlowAdapter) GetProvider() Provider {
	return Provider("siliconflow")
}

func (a *SiliconFlowAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	// 硅基流动使用 OpenAI 兼容的 API
	sfReq := a.convertToOpenAIFormat(req)
	
	reqBody, err := json.Marshal(sfReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("siliconflow api error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("siliconflow api error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var openaiResp models.ChatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &openaiResp, nil
}

func (a *SiliconFlowAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	sfReq := a.convertToOpenAIFormat(req)
	sfReq["stream"] = true
	
	reqBody, err := json.Marshal(sfReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("siliconflow stream error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("siliconflow stream error: status %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func (a *SiliconFlowAdapter) convertToOpenAIFormat(req *models.ChatCompletionRequest) map[string]interface{} {
	messages := make([]map[string]interface{}, 0, len(req.Messages))
	for _, msg := range req.Messages {
		msgMap := map[string]interface{}{
			"role": msg.Role,
		}
		
		switch v := msg.Content.(type) {
		case string:
			msgMap["content"] = v
		case []interface{}:
			msgMap["content"] = v
		default:
			msgMap["content"] = fmt.Sprintf("%v", v)
		}
		
		messages = append(messages, msgMap)
	}

	result := map[string]interface{}{
		"model":    req.Model,
		"messages": messages,
	}

	if req.Temperature != nil {
		result["temperature"] = *req.Temperature
	}
	if req.TopP != nil {
		result["top_p"] = *req.TopP
	}
	if req.MaxTokens != nil {
		result["max_tokens"] = *req.MaxTokens
	}
	if req.Stop != nil && len(req.Stop) > 0 {
		result["stop"] = req.Stop
	}

	return result
}


