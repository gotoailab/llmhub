package define

type Model string

// 具体模型厂商提供的模型

const (
	// OpenAI 下的模型
	ModelOpenAIGPT4          Model = "gpt-4"
	ModelOpenAIGPT4o         Model = "gpt-4o"
	ModelOpenAIGPT4oMini     Model = "gpt-4o-mini"
	ModelOpenAIGPT4Turbo     Model = "gpt-4-turbo"
	ModelOpenAIGPT4Turbo0125 Model = "gpt-4-turbo-preview"
	ModelOpenAIGPT35Turbo    Model = "gpt-3.5-turbo"
	ModelOpenAIGPT35Turbo16K Model = "gpt-3.5-turbo-16k"

	// Anthropic Claude 下的模型
	ModelClaude35Sonnet Model = "claude-3-5-sonnet-20241022"
	ModelClaude35Haiku  Model = "claude-3-5-haiku-20241022"
	ModelClaude3Opus    Model = "claude-3-opus-20240229"
	ModelClaude3Sonnet  Model = "claude-3-sonnet-20240229"
	ModelClaude3Haiku   Model = "claude-3-haiku-20240307"

	// Google Gemini 下的模型
	ModelGeminiPro       Model = "gemini-pro"
	ModelGeminiUltra     Model = "gemini-ultra"
	ModelGemini15Pro     Model = "gemini-1.5-pro"
	ModelGemini15Flash   Model = "gemini-1.5-flash"
	ModelGeminiProVision Model = "gemini-pro-vision"

	// Mistral AI 下的模型
	ModelMistralLarge  Model = "mistral-large"
	ModelMistralMedium Model = "mistral-medium"
	ModelMistralSmall  Model = "mistral-small"
	ModelMistral7B     Model = "mistral-7b"
	ModelMixtral8x7B   Model = "mixtral-8x7b"

	// DeepSeek 下的模型
	ModelDeepSeekV3      Model = "deepseek-chat"
	ModelDeepSeekV31     Model = "deepseek-v3.1"
	ModelDeepSeekV32Exp  Model = "deepseek-v3.2-exp"
	ModelDeepSeekR1      Model = "deepseek-r1"
	ModelDeepSeekCoder   Model = "deepseek-coder"
	ModelDeepSeekCoderV2 Model = "deepseek-coder-v2"

	// Groq 下的模型
	ModelGroqLlama31_70B Model = "llama-3.1-70b-versatile"
	ModelGroqLlama31_8B  Model = "llama-3.1-8b-instant"
	ModelGroqMixtral8x7B Model = "mixtral-8x7b-32768"
	ModelGroqGemma7B     Model = "gemma-7b-it"
	ModelGroqGemma2B     Model = "gemma-2b-it"

	// Cohere 下的模型
	ModelCohereCommandRPlus Model = "command-r-plus"
	ModelCohereCommandR     Model = "command-r"
	ModelCohereCommand      Model = "command"
	ModelCohereCommandLight Model = "command-light"

	// xAI 下的模型
	ModelXAIGrokBeta Model = "grok-beta"
	ModelXAIGrok2    Model = "grok-2"

	// Together.ai 下的模型
	ModelTogetherLlama270B   Model = "meta-llama/Llama-2-70b-chat-hf"
	ModelTogetherMixtral8x7B Model = "mistralai/Mixtral-8x7B-Instruct-v0.1"
	ModelTogetherQwen25_72B  Model = "Qwen/Qwen2.5-72B-Instruct"

	// Novita.ai 下的模型（平台模型，具体模型名可能变化）
	ModelNovitaDefault Model = "novita-default"

	// 通义千问 (Qwen) 下的模型
	ModelQwenTurbo Model = "qwen-turbo"
	ModelQwenPlus  Model = "qwen-plus"
	ModelQwenMax   Model = "qwen-max"
	ModelQwen72B   Model = "qwen-72b-chat"
	ModelQwen14B   Model = "qwen-14b-chat"
	ModelQwen7B    Model = "qwen-7b-chat"

	// 硅基流动 (SiliconFlow) 下的模型
	ModelSiliconFlowDeepSeekR1 Model = "DeepSeek-AI/DeepSeek-R1"
	ModelSiliconFlowDeepSeekV3 Model = "DeepSeek-AI/DeepSeek-V3"
	ModelSiliconFlowQwQ32B     Model = "Qwen/QwQ-32B"

	// 豆包 (Doubao) 下的模型
	ModelDoubaoPro    Model = "doubao-pro"
	ModelDoubaoLite   Model = "doubao-lite"
	ModelDoubaoPro4K  Model = "doubao-pro-4k"
	ModelDoubaoPro32K Model = "doubao-pro-32k"

	// 文心一言 (Ernie) 下的模型
	ModelErnie40       Model = "ernie-4.0"
	ModelErnie35       Model = "ernie-3.5"
	ModelErnieTurbo    Model = "ernie-turbo"
	ModelErnieBot      Model = "ernie-bot"
	ModelErnieBotTurbo Model = "ernie-bot-turbo"

	// 讯飞星火 (Spark) 下的模型
	ModelSparkV4   Model = "spark-v4"
	ModelSparkV35  Model = "spark-v3.5"
	ModelSparkLite Model = "spark-lite"
	ModelSparkMax  Model = "spark-max"

	// ChatGLM (智谱) 下的模型
	ModelGLM46     Model = "glm-4.6"
	ModelGLM45     Model = "glm-4.5"
	ModelGLM45Air  Model = "glm-4.5-air"
	ModelGLM4      Model = "glm-4"
	ModelGLM3Turbo Model = "glm-3-turbo"

	// 360智脑 下的模型
	Model360Brain Model = "360-brain"
	Model360GPT   Model = "360-gpt"

	// 腾讯混元 (Hunyuan) 下的模型
	ModelHunyuanPro      Model = "hunyuan-pro"
	ModelHunyuanStandard Model = "hunyuan-standard"
	ModelHunyuanLite     Model = "hunyuan-lite"

	// Moonshot AI (Kimi) 下的模型
	ModelMoonshotKimiK2Instruct Model = "moonshot-v1-8k"
	ModelMoonshotKimiK1         Model = "moonshot-v1-32k"
	ModelMoonshotKimiK15        Model = "moonshot-v1-128k"

	// 百川 (Baichuan) 下的模型
	ModelBaichuan2Turbo Model = "baichuan2-turbo"
	ModelBaichuan2_13B  Model = "baichuan2-13b-chat"
	ModelBaichuan2_7B   Model = "baichuan2-7b-chat"

	// MINIMAX 下的模型
	ModelMiniMaxAbab65gChat Model = "abab6.5g-chat"
	ModelMiniMaxAbab65tChat Model = "abab6.5t-chat"
	ModelMiniMaxAbab65sChat Model = "abab6.5s-chat"

	// 零一万物 (Yi) 下的模型
	ModelYi34B Model = "yi-34b-chat"
	ModelYi6B  Model = "yi-6b-chat"
	ModelYiVL  Model = "yi-vl"

	// 阶跃星辰 (StepFun) 下的模型
	ModelStepFunStep1 Model = "step-1"

	// Coze 下的模型（平台模型，具体模型名可能变化）
	ModelCozeDefault Model = "coze-default"

	// Ollama 下的模型（本地部署，模型名由用户自定义）
	ModelOllamaLlama2     Model = "llama2"
	ModelOllamaLlama3     Model = "llama3"
	ModelOllamaMistral    Model = "mistral"
	ModelOllamaMixtral    Model = "mixtral"
	ModelOllamaQwen       Model = "qwen"
	ModelOllamaGemma      Model = "gemma"
	ModelOllamaNeuralChat Model = "neural-chat"
	ModelOllamaStarCoder  Model = "starcoder"
	ModelOllamaCodeLlama  Model = "codellama"
	ModelOllamaPhi        Model = "phi"
)
