package aihub

import (
	"github.com/aihub/internal/adapters"
)

func init() {
	// 注册所有适配器
	adapters.Register("openai", adapters.NewOpenAIAdapter)
	adapters.Register("claude", adapters.NewClaudeAdapter)
	adapters.Register("deepseek", adapters.NewDeepSeekAdapter)
	adapters.Register("qwen", adapters.NewQwenAdapter)
	adapters.Register("siliconflow", adapters.NewSiliconFlowAdapter)
	adapters.Register("gemini", adapters.NewGeminiAdapter)
	adapters.Register("mistral", adapters.NewMistralAdapter)

	// 注册新模型适配器
	adapters.Register("doubao", adapters.NewDoubaoAdapter)
	adapters.Register("ernie", adapters.NewErnieAdapter)
	adapters.Register("spark", adapters.NewSparkAdapter)
	adapters.Register("chatglm", adapters.NewChatglmAdapter)
	adapters.Register("360", adapters.New360Adapter)
	adapters.Register("hunyuan", adapters.NewHunyuanAdapter)
	adapters.Register("moonshot", adapters.NewMoonshotAdapter)
	adapters.Register("baichuan", adapters.NewBaichuanAdapter)
	adapters.Register("minimax", adapters.NewMinimaxAdapter)
	adapters.Register("groq", adapters.NewGroqAdapter)
	adapters.Register("ollama", adapters.NewOllamaAdapter)
	adapters.Register("yi", adapters.NewYiAdapter)
	adapters.Register("stepfun", adapters.NewStepfunAdapter)
	adapters.Register("coze", adapters.NewCozeAdapter)
	adapters.Register("cohere", adapters.NewCohereAdapter)
	adapters.Register("together", adapters.NewTogetherAdapter)
	adapters.Register("novita", adapters.NewNovitaAdapter)
	adapters.Register("xai", adapters.NewXaiAdapter)
}
