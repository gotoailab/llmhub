package main

import (
	"fmt"

	"github.com/gotoailab/llmhub"
)

func main() {
	fmt.Println("支持的模型提供商:")
	fmt.Println("================")

	// 获取所有提供商
	providers := llmhub.AllProviders()

	// 按类别分组显示
	international := []llmhub.Provider{
		llmhub.ProviderOpenAI,
		llmhub.ProviderClaude,
		llmhub.ProviderGemini,
		llmhub.ProviderMistral,
		llmhub.ProviderDeepSeek,
		llmhub.ProviderGroq,
		llmhub.ProviderCohere,
		llmhub.ProviderXAI,
		llmhub.ProviderTogether,
		llmhub.ProviderNovita,
	}

	domestic := []llmhub.Provider{
		llmhub.ProviderQwen,
		llmhub.ProviderSiliconFlow,
		llmhub.ProviderDoubao,
		llmhub.ProviderErnie,
		llmhub.ProviderSpark,
		llmhub.ProviderChatGLM,
		llmhub.Provider360,
		llmhub.ProviderHunyuan,
		llmhub.ProviderMoonshot,
		llmhub.ProviderBaichuan,
		llmhub.ProviderMiniMax,
		llmhub.ProviderYi,
		llmhub.ProviderStepFun,
		llmhub.ProviderCoze,
	}

	local := []llmhub.Provider{
		llmhub.ProviderOllama,
	}

	fmt.Println("\n国际模型:")
	for _, p := range international {
		if contains(providers, p) {
			fmt.Printf("  - %s\n", p)
		}
	}

	fmt.Println("\n国内模型:")
	for _, p := range domestic {
		if contains(providers, p) {
			fmt.Printf("  - %s\n", p)
		}
	}

	fmt.Println("\n本地模型:")
	for _, p := range local {
		if contains(providers, p) {
			fmt.Printf("  - %s\n", p)
		}
	}

	fmt.Printf("\n总计: %d 个提供商\n", len(providers))
}

func contains(list []llmhub.Provider, item llmhub.Provider) bool {
	for _, p := range list {
		if p == item {
			return true
		}
	}
	return false
}
