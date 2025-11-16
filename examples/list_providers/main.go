package main

import (
	"fmt"

	"github.com/aihub"
)

func main() {
	fmt.Println("支持的模型提供商:")
	fmt.Println("================")

	// 获取所有提供商
	providers := aihub.AllProviders()

	// 按类别分组显示
	international := []aihub.Provider{
		aihub.ProviderOpenAI,
		aihub.ProviderClaude,
		aihub.ProviderGemini,
		aihub.ProviderMistral,
		aihub.ProviderDeepSeek,
		aihub.ProviderGroq,
		aihub.ProviderCohere,
		aihub.ProviderXAI,
		aihub.ProviderTogether,
		aihub.ProviderNovita,
	}

	domestic := []aihub.Provider{
		aihub.ProviderQwen,
		aihub.ProviderSiliconFlow,
		aihub.ProviderDoubao,
		aihub.ProviderErnie,
		aihub.ProviderSpark,
		aihub.ProviderChatGLM,
		aihub.Provider360,
		aihub.ProviderHunyuan,
		aihub.ProviderMoonshot,
		aihub.ProviderBaichuan,
		aihub.ProviderMiniMax,
		aihub.ProviderYi,
		aihub.ProviderStepFun,
		aihub.ProviderCoze,
	}

	local := []aihub.Provider{
		aihub.ProviderOllama,
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

func contains(list []aihub.Provider, item aihub.Provider) bool {
	for _, p := range list {
		if p == item {
			return true
		}
	}
	return false
}

