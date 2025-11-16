package llmhub

// Provider 支持的模型提供商枚举
type Provider string

const (
	// 国际模型
	ProviderOpenAI   Provider = "openai"   // OpenAI
	ProviderClaude   Provider = "claude"   // Anthropic Claude
	ProviderGemini   Provider = "gemini"   // Google Gemini / PaLM2
	ProviderMistral  Provider = "mistral"  // Mistral AI
	ProviderDeepSeek Provider = "deepseek" // DeepSeek
	ProviderGroq     Provider = "groq"     // Groq
	ProviderCohere   Provider = "cohere"   // Cohere
	ProviderXAI      Provider = "xai"      // xAI
	ProviderTogether Provider = "together" // together.ai
	ProviderNovita   Provider = "novita"   // novita.ai

	// 国内模型
	ProviderQwen        Provider = "qwen"        // 通义千问
	ProviderSiliconFlow Provider = "siliconflow" // 硅基流动
	ProviderDoubao      Provider = "doubao"      // 豆包（字节跳动）
	ProviderErnie       Provider = "ernie"       // 文心一言（百度）
	ProviderSpark       Provider = "spark"       // 讯飞星火
	ProviderChatGLM     Provider = "chatglm"     // ChatGLM（智谱）
	Provider360         Provider = "360"         // 360智脑
	ProviderHunyuan     Provider = "hunyuan"     // 腾讯混元
	ProviderMoonshot    Provider = "moonshot"    // Moonshot AI
	ProviderBaichuan    Provider = "baichuan"    // 百川大模型
	ProviderMiniMax     Provider = "minimax"     // MINIMAX
	ProviderYi          Provider = "yi"          // 零一万物
	ProviderStepFun     Provider = "stepfun"     // 阶跃星辰
	ProviderCoze        Provider = "coze"        // Coze

	// 本地模型
	ProviderOllama Provider = "ollama" // Ollama
)

// String 返回提供商的字符串表示
func (p Provider) String() string {
	return string(p)
}

// IsValid 检查 Provider 是否有效
func (p Provider) IsValid() bool {
	switch p {
	case ProviderOpenAI, ProviderClaude, ProviderGemini, ProviderMistral,
		ProviderDeepSeek, ProviderGroq, ProviderCohere, ProviderXAI,
		ProviderTogether, ProviderNovita, ProviderQwen, ProviderSiliconFlow,
		ProviderDoubao, ProviderErnie, ProviderSpark, ProviderChatGLM,
		Provider360, ProviderHunyuan, ProviderMoonshot, ProviderBaichuan,
		ProviderMiniMax, ProviderYi, ProviderStepFun, ProviderCoze,
		ProviderOllama:
		return true
	default:
		return false
	}
}

// AllProviders 返回所有支持的提供商列表
func AllProviders() []Provider {
	return []Provider{
		ProviderOpenAI,
		ProviderClaude,
		ProviderGemini,
		ProviderMistral,
		ProviderDeepSeek,
		ProviderGroq,
		ProviderCohere,
		ProviderXAI,
		ProviderTogether,
		ProviderNovita,
		ProviderQwen,
		ProviderSiliconFlow,
		ProviderDoubao,
		ProviderErnie,
		ProviderSpark,
		ProviderChatGLM,
		Provider360,
		ProviderHunyuan,
		ProviderMoonshot,
		ProviderBaichuan,
		ProviderMiniMax,
		ProviderYi,
		ProviderStepFun,
		ProviderCoze,
		ProviderOllama,
	}
}
