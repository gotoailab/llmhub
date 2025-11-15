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

type ClaudeAdapter struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewClaudeAdapter 创建 Claude 适配器（导出以供注册）
func NewClaudeAdapter(apiKey, baseURL string) (Adapter, error) {
	if baseURL == "" {
		baseURL = "https://api.anthropic.com/v1"
	}
	
	return &ClaudeAdapter{
		apiKey:  apiKey,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}, nil
}

func (a *ClaudeAdapter) GetProvider() string {
	return "claude"
}

func (a *ClaudeAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	// Claude API 格式转换
	claudeReq := a.convertToClaudeRequest(req)
	
	reqBody, err := json.Marshal(claudeReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/messages", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", a.apiKey)
	httpReq.Header.Set("anthropic-version", "2023-06-01")

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("claude api error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("claude api error: status %d, body: %s", resp.StatusCode, string(body))
	}

	var claudeResp ClaudeResponse
	if err := json.NewDecoder(resp.Body).Decode(&claudeResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// 转换回 OpenAI 格式
	return a.convertFromClaudeResponse(&claudeResp, req.Model), nil
}

func (a *ClaudeAdapter) ChatCompletionStream(ctx context.Context, req *models.ChatCompletionRequest) (io.ReadCloser, error) {
	// Claude 流式响应实现
	claudeReq := a.convertToClaudeRequest(req)
	claudeReq.Stream = true
	
	reqBody, err := json.Marshal(claudeReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/messages", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", a.apiKey)
	httpReq.Header.Set("anthropic-version", "2023-06-01")

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("claude stream error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("claude stream error: status %d", resp.StatusCode)
	}

	return resp.Body, nil
}

// Claude 请求结构
type ClaudeRequest struct {
	Model       string          `json:"model"`
	MaxTokens   int            `json:"max_tokens"`
	Messages    []ClaudeMessage `json:"messages"`
	Temperature *float64        `json:"temperature,omitempty"`
	TopP        *float64        `json:"top_p,omitempty"`
	Stream      bool            `json:"stream,omitempty"`
}

type ClaudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Claude 响应结构
type ClaudeResponse struct {
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	Role         string          `json:"role"`
	Content      []ClaudeContent `json:"content"`
	Model        string          `json:"model"`
	StopReason   string          `json:"stop_reason"`
	StopSequence string          `json:"stop_sequence,omitempty"`
	Usage        ClaudeUsage     `json:"usage"`
}

type ClaudeContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ClaudeUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

func (a *ClaudeAdapter) convertToClaudeRequest(req *models.ChatCompletionRequest) *ClaudeRequest {
	messages := make([]ClaudeMessage, 0, len(req.Messages))
	for _, msg := range req.Messages {
		var content string
		switch v := msg.Content.(type) {
		case string:
			content = v
		default:
			content = fmt.Sprintf("%v", v)
		}
		
		// Claude 使用 "user" 和 "assistant"，需要转换
		role := msg.Role
		if role == "system" {
			// Claude 不支持 system role，可以合并到第一个 user message
			continue
		}
		
		messages = append(messages, ClaudeMessage{
			Role:    role,
			Content: content,
		})
	}

	claudeReq := &ClaudeRequest{
		Model:     a.mapModelName(req.Model),
		Messages:  messages,
		MaxTokens: getIntValue(req.MaxTokens),
		Temperature: req.Temperature,
		TopP:      req.TopP,
	}

	if claudeReq.MaxTokens == 0 {
		claudeReq.MaxTokens = 4096 // 默认值
	}

	return claudeReq
}

func (a *ClaudeAdapter) convertFromClaudeResponse(claudeResp *ClaudeResponse, modelName string) *models.ChatCompletionResponse {
	var content string
	if len(claudeResp.Content) > 0 {
		content = claudeResp.Content[0].Text
	}

	return &models.ChatCompletionResponse{
		ID:      claudeResp.ID,
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   modelName,
		Choices: []models.ChatCompletionChoice{
			{
				Index: 0,
				Message: models.ChatMessage{
					Role:    "assistant",
					Content: content,
				},
				FinishReason: claudeResp.StopReason,
			},
		},
		Usage: models.Usage{
			PromptTokens:     claudeResp.Usage.InputTokens,
			CompletionTokens: claudeResp.Usage.OutputTokens,
			TotalTokens:      claudeResp.Usage.InputTokens + claudeResp.Usage.OutputTokens,
		},
	}
}

func (a *ClaudeAdapter) mapModelName(modelName string) string {
	// 模型名称映射
	mapping := map[string]string{
		"claude-3-opus":    "claude-3-opus-20240229",
		"claude-3-sonnet":  "claude-3-sonnet-20240229",
		"claude-3-haiku":   "claude-3-haiku-20240307",
		"claude-3-5-sonnet": "claude-3-5-sonnet-20241022",
	}
	
	if mapped, ok := mapping[modelName]; ok {
		return mapped
	}
	return modelName
}


