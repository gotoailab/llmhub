package aihub

import (
	"github.com/aihub/internal/adapters"
)

func init() {
	// 注册所有适配器（使用枚举，显式转换为 adapters.Provider）
	adapters.Register(adapters.Provider(ProviderOpenAI), adapters.NewOpenAIAdapter)
	adapters.Register(adapters.Provider(ProviderClaude), adapters.NewClaudeAdapter)
	adapters.Register(adapters.Provider(ProviderDeepSeek), adapters.NewDeepSeekAdapter)
	adapters.Register(adapters.Provider(ProviderQwen), adapters.NewQwenAdapter)
	adapters.Register(adapters.Provider(ProviderSiliconFlow), adapters.NewSiliconFlowAdapter)
	adapters.Register(adapters.Provider(ProviderGemini), adapters.NewGeminiAdapter)
	adapters.Register(adapters.Provider(ProviderMistral), adapters.NewMistralAdapter)

	// 注册新模型适配器
	adapters.Register(adapters.Provider(ProviderDoubao), adapters.NewDoubaoAdapter)
	adapters.Register(adapters.Provider(ProviderErnie), adapters.NewErnieAdapter)
	adapters.Register(adapters.Provider(ProviderSpark), adapters.NewSparkAdapter)
	adapters.Register(adapters.Provider(ProviderChatGLM), adapters.NewChatglmAdapter)
	adapters.Register(adapters.Provider(Provider360), adapters.New360Adapter)
	adapters.Register(adapters.Provider(ProviderHunyuan), adapters.NewHunyuanAdapter)
	adapters.Register(adapters.Provider(ProviderMoonshot), adapters.NewMoonshotAdapter)
	adapters.Register(adapters.Provider(ProviderBaichuan), adapters.NewBaichuanAdapter)
	adapters.Register(adapters.Provider(ProviderMiniMax), adapters.NewMinimaxAdapter)
	adapters.Register(adapters.Provider(ProviderGroq), adapters.NewGroqAdapter)
	adapters.Register(adapters.Provider(ProviderOllama), adapters.NewOllamaAdapter)
	adapters.Register(adapters.Provider(ProviderYi), adapters.NewYiAdapter)
	adapters.Register(adapters.Provider(ProviderStepFun), adapters.NewStepfunAdapter)
	adapters.Register(adapters.Provider(ProviderCoze), adapters.NewCozeAdapter)
	adapters.Register(adapters.Provider(ProviderCohere), adapters.NewCohereAdapter)
	adapters.Register(adapters.Provider(ProviderTogether), adapters.NewTogetherAdapter)
	adapters.Register(adapters.Provider(ProviderNovita), adapters.NewNovitaAdapter)
	adapters.Register(adapters.Provider(ProviderXAI), adapters.NewXaiAdapter)
}
