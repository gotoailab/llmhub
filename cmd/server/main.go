package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gotoailab/llmhub"
	"github.com/gotoailab/llmhub/internal/adapters"
	"github.com/gotoailab/llmhub/internal/api"
	"github.com/gotoailab/llmhub/internal/config"
)

func init() {
	// 注册所有适配器（使用枚举）
	adapters.Register(adapters.Provider(llmhub.ProviderOpenAI), adapters.NewOpenAIAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderClaude), adapters.NewClaudeAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderDeepSeek), adapters.NewDeepSeekAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderQwen), adapters.NewQwenAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderSiliconFlow), adapters.NewSiliconFlowAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderGemini), adapters.NewGeminiAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderMistral), adapters.NewMistralAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderDoubao), adapters.NewDoubaoAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderErnie), adapters.NewErnieAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderSpark), adapters.NewSparkAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderChatGLM), adapters.NewChatglmAdapter)
	adapters.Register(adapters.Provider(llmhub.Provider360), adapters.New360Adapter)
	adapters.Register(adapters.Provider(llmhub.ProviderHunyuan), adapters.NewHunyuanAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderMoonshot), adapters.NewMoonshotAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderBaichuan), adapters.NewBaichuanAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderMiniMax), adapters.NewMinimaxAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderGroq), adapters.NewGroqAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderOllama), adapters.NewOllamaAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderYi), adapters.NewYiAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderStepFun), adapters.NewStepfunAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderCoze), adapters.NewCozeAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderCohere), adapters.NewCohereAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderTogether), adapters.NewTogetherAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderNovita), adapters.NewNovitaAdapter)
	adapters.Register(adapters.Provider(llmhub.ProviderXAI), adapters.NewXaiAdapter)
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.Parse()

	// 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 创建处理器和路由
	handler := api.NewHandler()
	router := api.SetupRouter(handler)

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting server on %s", addr)
	log.Printf("Loaded %d models", len(cfg.Models))

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
