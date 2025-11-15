package models

// OpenAI 兼容的请求结构
type ChatCompletionRequest struct {
	Model            string                 `json:"model"`
	Messages         []ChatMessage          `json:"messages"`
	Temperature      *float64               `json:"temperature,omitempty"`
	TopP             *float64               `json:"top_p,omitempty"`
	MaxTokens        *int                   `json:"max_tokens,omitempty"`
	Stream           bool                   `json:"stream,omitempty"`
	PresencePenalty  *float64               `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64               `json:"frequency_penalty,omitempty"`
	Stop             []string               `json:"stop,omitempty"`
	User             string                 `json:"user,omitempty"`
	Functions        []FunctionDefinition  `json:"functions,omitempty"`
	FunctionCall     interface{}            `json:"function_call,omitempty"`
	LogitBias       map[string]int         `json:"logit_bias,omitempty"`
	LogProbs         bool                   `json:"logprobs,omitempty"`
	TopLogProbs      *int                   `json:"top_logprobs,omitempty"`
	ResponseFormat   *ResponseFormat        `json:"response_format,omitempty"`
	Seed             *int                   `json:"seed,omitempty"`
	Tools            []Tool                 `json:"tools,omitempty"`
	ToolChoice       interface{}            `json:"tool_choice,omitempty"`
	ExtraParams      map[string]interface{} `json:"-"`
}

type ChatMessage struct {
	Role         string          `json:"role"`
	Content      interface{}     `json:"content"`
	Name         string          `json:"name,omitempty"`
	FunctionCall *FunctionCall   `json:"function_call,omitempty"`
	ToolCalls    []ToolCall      `json:"tool_calls,omitempty"`
	ToolCallID   string          `json:"tool_call_id,omitempty"`
}

type FunctionDefinition struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Parameters  interface{} `json:"parameters,omitempty"`
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ToolCall struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Function FunctionCall `json:"function"`
}

type Tool struct {
	Type     string           `json:"type"`
	Function FunctionDefinition `json:"function"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

