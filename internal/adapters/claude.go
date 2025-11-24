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

func (a *ClaudeAdapter) GetProvider() Provider {
	return Provider("claude")
}

func (a *ClaudeAdapter) ChatCompletion(ctx context.Context, req *models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	// 检查是否有工具调用但模型不支持
	if len(req.Tools) > 0 || len(req.Functions) > 0 {
		// Claude 3.5 Sonnet 和更新版本支持工具调用
		if !a.supportsTools(req.Model) {
			return nil, fmt.Errorf("tool use not supported for model %s", req.Model)
		}
	}

	// Claude API 格式转换
	claudeReq, err := a.convertToClaudeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert request: %w", err)
	}

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
	// 检查工具支持
	if (len(req.Tools) > 0 || len(req.Functions) > 0) && !a.supportsTools(req.Model) {
		return nil, fmt.Errorf("tool use not supported for model %s", req.Model)
	}

	// Claude 流式响应实现
	claudeReq, err := a.convertToClaudeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert request: %w", err)
	}
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

// 检查模型是否支持工具调用
func (a *ClaudeAdapter) supportsTools(modelName string) bool {
	supportedModels := map[string]bool{
		"claude-3-5-sonnet":           true,
		"claude-3-5-sonnet-20241022":  true,
		"claude-3-opus":               true,
		"claude-3-opus-20240229":      true,
		"claude-3-sonnet":             true,
		"claude-3-sonnet-20240229":    true,
		"claude-3-haiku":              false, // Haiku 不支持工具调用
		"claude-3-haiku-20240307":     false,
	}

	supported, exists := supportedModels[modelName]
	if !exists {
		// 对于未知模型，假设新版本支持工具调用
		return true
	}
	return supported
}

// Claude 请求结构
type ClaudeRequest struct {
	Model       string          `json:"model"`
	MaxTokens   int             `json:"max_tokens"`
	Messages    []ClaudeMessage `json:"messages"`
	Temperature *float64        `json:"temperature,omitempty"`
	TopP        *float64        `json:"top_p,omitempty"`
	Stream      bool            `json:"stream,omitempty"`
	Tools       []ClaudeTool    `json:"tools,omitempty"`
	ToolChoice  interface{}     `json:"tool_choice,omitempty"`
	System      string          `json:"system,omitempty"`
}

type ClaudeMessage struct {
	Role    string        `json:"role"`
	Content interface{}   `json:"content"`
}

type ClaudeTool struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	InputSchema interface{} `json:"input_schema,omitempty"`
}

type ClaudeToolUse struct {
	Type  string      `json:"type"`
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Input interface{} `json:"input"`
}

type ClaudeToolResult struct {
	Type       string      `json:"type"`
	ToolUseID  string      `json:"tool_use_id"`
	Content    interface{} `json:"content"`
	IsError    bool        `json:"is_error,omitempty"`
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
	Type  string      `json:"type"`
	Text  string      `json:"text,omitempty"`
	ID    string      `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Input interface{} `json:"input,omitempty"`
}

type ClaudeUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

func (a *ClaudeAdapter) convertToClaudeRequest(req *models.ChatCompletionRequest) (*ClaudeRequest, error) {
	messages := make([]ClaudeMessage, 0, len(req.Messages))
	var systemMessage string

	for _, msg := range req.Messages {
		// 处理 system 消息
		if msg.Role == "system" {
			if content, ok := msg.Content.(string); ok {
				systemMessage = content
			}
			continue
		}

		claudeMsg := ClaudeMessage{
			Role: msg.Role,
		}

		// 处理工具调用消息
		if len(msg.ToolCalls) > 0 {
			content := make([]interface{}, 0)
			
			// 添加文本内容（如果有）
			if textContent, ok := msg.Content.(string); ok && textContent != "" {
				content = append(content, map[string]string{
					"type": "text",
					"text": textContent,
				})
			}

			// 添加工具调用
			for _, tc := range msg.ToolCalls {
				var input interface{}
				if tc.Function.Arguments != "" {
					json.Unmarshal([]byte(tc.Function.Arguments), &input)
				}
				
				content = append(content, ClaudeToolUse{
					Type:  "tool_use",
					ID:    tc.ID,
					Name:  tc.Function.Name,
					Input: input,
				})
			}
			claudeMsg.Content = content
		} else if msg.ToolCallID != "" {
			// 处理工具结果消息
			claudeMsg.Content = []ClaudeToolResult{
				{
					Type:      "tool_result",
					ToolUseID: msg.ToolCallID,
					Content:   msg.Content,
				},
			}
		} else {
			// 普通文本消息
			claudeMsg.Content = msg.Content
		}

		messages = append(messages, claudeMsg)
	}

	claudeReq := &ClaudeRequest{
		Model:       a.mapModelName(req.Model),
		Messages:    messages,
		MaxTokens:   getIntValue(req.MaxTokens),
		Temperature: req.Temperature,
		TopP:        req.TopP,
		System:      systemMessage,
	}

	if claudeReq.MaxTokens == 0 {
		claudeReq.MaxTokens = 4096 // 默认值
	}

	// 转换工具定义
	if len(req.Tools) > 0 {
		tools := make([]ClaudeTool, 0, len(req.Tools))
		for _, tool := range req.Tools {
			tools = append(tools, ClaudeTool{
				Name:        tool.Function.Name,
				Description: tool.Function.Description,
				InputSchema: tool.Function.Parameters,
			})
		}
		claudeReq.Tools = tools

		if req.ToolChoice != nil {
			claudeReq.ToolChoice = req.ToolChoice
		}
	}

	// 兼容旧的 Functions 格式
	if len(req.Functions) > 0 {
		tools := make([]ClaudeTool, 0, len(req.Functions))
		for _, fn := range req.Functions {
			tools = append(tools, ClaudeTool{
				Name:        fn.Name,
				Description: fn.Description,
				InputSchema: fn.Parameters,
			})
		}
		claudeReq.Tools = tools
	}

	return claudeReq, nil
}

func (a *ClaudeAdapter) convertFromClaudeResponse(claudeResp *ClaudeResponse, modelName string) *models.ChatCompletionResponse {
	message := models.ChatMessage{
		Role: "assistant",
	}

	var textContent string
	var toolCalls []models.ToolCall

	// 处理响应内容
	for _, content := range claudeResp.Content {
		switch content.Type {
		case "text":
			textContent = content.Text
		case "tool_use":
			// 将输入参数序列化为 JSON 字符串
			args := ""
			if content.Input != nil {
				if argsBytes, err := json.Marshal(content.Input); err == nil {
					args = string(argsBytes)
				}
			}

			toolCalls = append(toolCalls, models.ToolCall{
				ID:   content.ID,
				Type: "function",
				Function: models.FunctionCall{
					Name:      content.Name,
					Arguments: args,
				},
			})
		}
	}

	message.Content = textContent
	if len(toolCalls) > 0 {
		message.ToolCalls = toolCalls
	}

	return &models.ChatCompletionResponse{
		ID:      claudeResp.ID,
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   modelName,
		Choices: []models.ChatCompletionChoice{
			{
				Index:        0,
				Message:      message,
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
		"claude-3-opus":     "claude-3-opus-20240229",
		"claude-3-sonnet":   "claude-3-sonnet-20240229",
		"claude-3-haiku":    "claude-3-haiku-20240307",
		"claude-3-5-sonnet": "claude-3-5-sonnet-20241022",
	}

	if mapped, ok := mapping[modelName]; ok {
		return mapped
	}
	return modelName
}
