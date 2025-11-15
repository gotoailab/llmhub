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

type QwenAdapter struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewQwenAdapter 创建 Qwen 适配器（导出以供注册）
func NewQwenAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://dashscope.aliyuncs.com/api/v1"
	}
	
	return &QwenAdapter{
		apiKey:  apiKey,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}, nil
}

func (a *QwenAdapter) GetProvider() string {
	return "qwen"
}

func (a *QwenAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	// Qwen 使用 OpenAI 兼容的 API
	qwenReq := a.convertToOpenAIFormat(req)
	
	reqBody, err := json.Marshal(qwenReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/services/aigc/text-generation/generation", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)
	httpReq.Header.Set("X-DashScope-SSE", "disable")

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("qwen api error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("qwen api error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var qwenResp QwenResponse
	if err := json.NewDecoder(resp.Body).Decode(&qwenResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// 转换回 OpenAI 格式
	return a.convertFromQwenResponse(&qwenResp, req.Model), nil
}

func (a *QwenAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	qwenReq := a.convertToOpenAIFormat(req)
	qwenReq["stream"] = true
	
	reqBody, err := json.Marshal(qwenReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/services/aigc/text-generation/generation", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)
	httpReq.Header.Set("X-DashScope-SSE", "enable")

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("qwen stream error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("qwen stream error: status %d", resp.StatusCode)
	}

	return resp.Body, nil
}

type QwenResponse struct {
	RequestID string      `json:"request_id"`
	Output    QwenOutput  `json:"output"`
	Usage     QwenUsage   `json:"usage"`
}

type QwenOutput struct {
	Text         string `json:"text"`
	FinishReason string `json:"finish_reason"`
}

type QwenUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

func (a *QwenAdapter) convertToOpenAIFormat(req *models.ChatCompletionRequest) map[string]interface{} {
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
		"model":    a.mapModelName(req.Model),
		"input": map[string]interface{}{
			"messages": messages,
		},
	}

	if req.Temperature != nil {
		result["parameters"] = map[string]interface{}{
			"temperature": *req.Temperature,
		}
	}
	if req.TopP != nil {
		if params, ok := result["parameters"].(map[string]interface{}); ok {
			params["top_p"] = *req.TopP
		} else {
			result["parameters"] = map[string]interface{}{
				"top_p": *req.TopP,
			}
		}
	}
	if req.MaxTokens != nil {
		if params, ok := result["parameters"].(map[string]interface{}); ok {
			params["max_tokens"] = *req.MaxTokens
		} else {
			result["parameters"] = map[string]interface{}{
				"max_tokens": *req.MaxTokens,
			}
		}
	}

	return result
}

func (a *QwenAdapter) convertFromQwenResponse(qwenResp *QwenResponse, modelName string) *models.ChatCompletionResponse {
	return &models.ChatCompletionResponse{
		ID:      qwenResp.RequestID,
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   modelName,
		Choices: []models.ChatCompletionChoice{
			{
				Index: 0,
				Message: models.ChatMessage{
					Role:    "assistant",
					Content: qwenResp.Output.Text,
				},
				FinishReason: qwenResp.Output.FinishReason,
			},
		},
		Usage: models.Usage{
			PromptTokens:     qwenResp.Usage.InputTokens,
			CompletionTokens: qwenResp.Usage.OutputTokens,
			TotalTokens:      qwenResp.Usage.TotalTokens,
		},
	}
}

func (a *QwenAdapter) mapModelName(modelName string) string {
	// Qwen 模型名称映射
	mapping := map[string]string{
		"qwen-turbo":    "qwen-turbo",
		"qwen-plus":     "qwen-plus",
		"qwen-max":      "qwen-max",
		"qwen-max-long": "qwen-max-longcontext",
	}
	
	if mapped, ok := mapping[modelName]; ok {
		return mapped
	}
	return modelName
}


