package adapters

import (
	"context"
	"testing"

	"github.com/gotoailab/llmhub/internal/models"
)

func TestToolUseSupport(t *testing.T) {
	tests := []struct {
		name               string
		provider           Provider
		expectsToolSupport bool
	}{
		{"OpenAI", "openai", true},
		{"Claude", "claude", true},
		{"OpenRouter", "openrouter", true},
		{"Groq", "groq", true},
		{"DeepSeek", "deepseek", true},
		{"Together", "together", true},
		{"SiliconFlow", "siliconflow", true},
		{"Moonshot", "moonshot", true},
		{"StepFun", "stepfun", true},
		{"Mistral", "mistral", true},
		{"Cohere", "cohere", true},
		{"Novita", "novita", true},
		{"xAI", "xai", true},
		{"Qwen", "qwen", false},
		{"Baichuan", "baichuan", false},
		{"ChatGLM", "chatglm", false},
		{"Ernie", "ernie", false},
		{"Spark", "spark", false},
		{"Hunyuan", "hunyuan", false},
		{"360", "360", false},
		{"MiniMax", "minimax", false},
		{"Yi", "yi", false},
		{"Doubao", "doubao", false},
		{"Ollama", "ollama", false},
		{"Coze", "coze", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isProviderSupportsTools(tt.provider)
			if actual != tt.expectsToolSupport {
				t.Errorf("Provider %s: expected tool support %v, got %v", 
					tt.provider, tt.expectsToolSupport, actual)
			}
		})
	}
}

func TestConvertToOpenAIFormatGeneric(t *testing.T) {
	// 测试基本转换
	req := &models.ChatCompletionRequest{
		Model: "gpt-4",
		Messages: []models.ChatMessage{
			{Role: "user", Content: "Hello"},
		},
		Temperature: floatPtr(0.7),
		MaxTokens:   intPtr(100),
	}

	result := convertToOpenAIFormatGeneric(req)
	
	if result["model"] != "gpt-4" {
		t.Errorf("Expected model 'gpt-4', got %v", result["model"])
	}
	
	if result["temperature"] != 0.7 {
		t.Errorf("Expected temperature 0.7, got %v", result["temperature"])
	}
	
	if result["max_tokens"] != 100 {
		t.Errorf("Expected max_tokens 100, got %v", result["max_tokens"])
	}
}

func TestToolConversion(t *testing.T) {
	// 测试工具转换
	tools := []models.Tool{
		{
			Type: "function",
			Function: models.FunctionDefinition{
				Name:        "get_weather",
				Description: "Get weather info",
				Parameters: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"city": map[string]interface{}{
							"type": "string",
						},
					},
				},
			},
		},
	}

	req := &models.ChatCompletionRequest{
		Model:    "gpt-4",
		Messages: []models.ChatMessage{{Role: "user", Content: "What's the weather?"}},
		Tools:    tools,
		ToolChoice: "auto",
	}

	result := convertToOpenAIFormatGeneric(req)
	
	if _, ok := result["tools"]; !ok {
		t.Error("Expected tools field in converted request")
	}
	
	if result["tool_choice"] != "auto" {
		t.Errorf("Expected tool_choice 'auto', got %v", result["tool_choice"])
	}
}

func TestClaudeToolSupport(t *testing.T) {
	adapter := &ClaudeAdapter{}
	
	tests := []struct {
		model     string
		supported bool
	}{
		{"claude-3-5-sonnet", true},
		{"claude-3-opus", true},
		{"claude-3-sonnet", true},
		{"claude-3-haiku", false},
		{"unknown-model", true}, // 未知模型默认支持
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			actual := adapter.supportsTools(tt.model)
			if actual != tt.supported {
				t.Errorf("Model %s: expected tool support %v, got %v", 
					tt.model, tt.supported, actual)
			}
		})
	}
}

func TestToolUseErrorHandling(t *testing.T) {
	// 测试 Qwen 适配器的错误处理
	adapter := &QwenAdapter{
		apiKey:  "test-key",
		baseURL: "https://test.com",
	}

	req := &models.ChatCompletionRequest{
		Model: "qwen-turbo",
		Messages: []models.ChatMessage{
			{Role: "user", Content: "Test"},
		},
		Tools: []models.Tool{
			{
				Type: "function",
				Function: models.FunctionDefinition{
					Name: "test_tool",
				},
			},
		},
	}

	_, err := adapter.ChatCompletion(context.Background(), req)
	if err == nil {
		t.Error("Expected error for tool use with Qwen, got nil")
	}

	expectedMsg := "tool use not supported for Qwen models"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

// 辅助函数
func floatPtr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}