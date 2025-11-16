package llmhub

// ChatCompletionRequest OpenAI 兼容的请求结构
type ChatCompletionRequest struct {
	Model            string               `json:"model"`
	Messages         []ChatMessage        `json:"messages"`
	Temperature      *float64             `json:"temperature,omitempty"`
	TopP             *float64             `json:"top_p,omitempty"`
	MaxTokens        *int                 `json:"max_tokens,omitempty"`
	Stream           bool                 `json:"stream,omitempty"`
	PresencePenalty  *float64             `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64             `json:"frequency_penalty,omitempty"`
	Stop             []string             `json:"stop,omitempty"`
	User             string               `json:"user,omitempty"`
	Functions        []FunctionDefinition `json:"functions,omitempty"`
	FunctionCall     interface{}          `json:"function_call,omitempty"`
	LogitBias        map[string]int       `json:"logit_bias,omitempty"`
	LogProbs         bool                 `json:"logprobs,omitempty"`
	TopLogProbs      *int                 `json:"top_logprobs,omitempty"`
	ResponseFormat   *ResponseFormat      `json:"response_format,omitempty"`
	Seed             *int                 `json:"seed,omitempty"`
	Tools            []Tool               `json:"tools,omitempty"`
	ToolChoice       interface{}          `json:"tool_choice,omitempty"`
}

// ChatMessage 聊天消息
type ChatMessage struct {
	Role         string        `json:"role"`
	Content      interface{}   `json:"content"`
	Name         string        `json:"name,omitempty"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
	ToolCalls    []ToolCall    `json:"tool_calls,omitempty"`
	ToolCallID   string        `json:"tool_call_id,omitempty"`
}

// FunctionDefinition 函数定义
type FunctionDefinition struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Parameters  interface{} `json:"parameters,omitempty"`
}

// FunctionCall 函数调用
type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// ToolCall 工具调用
type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Function FunctionCall `json:"function"`
}

// Tool 工具
type Tool struct {
	Type     string             `json:"type"`
	Function FunctionDefinition `json:"function"`
}

// ResponseFormat 响应格式
type ResponseFormat struct {
	Type string `json:"type"`
}

// ChatCompletionResponse OpenAI 兼容的响应结构
type ChatCompletionResponse struct {
	ID                string                 `json:"id"`
	Object            string                 `json:"object"`
	Created           int64                  `json:"created"`
	Model             string                 `json:"model"`
	Choices           []ChatCompletionChoice `json:"choices"`
	Usage             Usage                  `json:"usage"`
	SystemFingerprint string                 `json:"system_fingerprint,omitempty"`
}

// ChatCompletionChoice 聊天完成选择
type ChatCompletionChoice struct {
	Index        int          `json:"index"`
	Message      ChatMessage  `json:"message"`
	FinishReason string       `json:"finish_reason"`
	Delta        *ChatMessage `json:"delta,omitempty"` // 用于流式响应
}

// Usage Token 使用情况
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail 错误详情
type ErrorDetail struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param,omitempty"`
	Code    string `json:"code,omitempty"`
}
