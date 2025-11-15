package aihub

// 内部类型定义，用于与 adapters 包交互

type internalChatCompletionRequest struct {
	Model            string
	Messages         []internalChatMessage
	Temperature      *float64
	TopP             *float64
	MaxTokens        *int
	Stream           bool
	PresencePenalty  *float64
	FrequencyPenalty *float64
	Stop             []string
	User             string
	Functions        []internalFunctionDefinition
	FunctionCall     interface{}
	LogitBias        map[string]int
	LogProbs         bool
	TopLogProbs      *int
	ResponseFormat   *internalResponseFormat
	Seed             *int
	Tools            []internalTool
	ToolChoice       interface{}
}

type internalChatMessage struct {
	Role         string
	Content      interface{}
	Name         string
	FunctionCall *internalFunctionCall
	ToolCalls    []internalToolCall
	ToolCallID   string
}

type internalFunctionDefinition struct {
	Name        string
	Description string
	Parameters  interface{}
}

type internalFunctionCall struct {
	Name      string
	Arguments string
}

type internalToolCall struct {
	ID       string
	Type     string
	Function internalFunctionCall
}

type internalTool struct {
	Type     string
	Function internalFunctionDefinition
}

type internalResponseFormat struct {
	Type string
}

type internalChatCompletionResponse struct {
	ID                string
	Object            string
	Created           int64
	Model             string
	Choices           []internalChatCompletionChoice
	Usage             internalUsage
	SystemFingerprint string
}

type internalChatCompletionChoice struct {
	Index        int
	Message      internalChatMessage
	FinishReason string
	Delta        *internalChatMessage
}

type internalUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

