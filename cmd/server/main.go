package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aihub"
	"github.com/aihub/internal/adapters"
	"github.com/aihub/internal/api"
	"github.com/aihub/internal/config"
)

func init() {
	// 注册所有适配器（使用枚举）
	adapters.Register(adapters.Provider(aihub.ProviderOpenAI), adapters.NewOpenAIAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderClaude), adapters.NewClaudeAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderDeepSeek), adapters.NewDeepSeekAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderQwen), adapters.NewQwenAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderSiliconFlow), adapters.NewSiliconFlowAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderGemini), adapters.NewGeminiAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderMistral), adapters.NewMistralAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderDoubao), adapters.NewDoubaoAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderErnie), adapters.NewErnieAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderSpark), adapters.NewSparkAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderChatGLM), adapters.NewChatglmAdapter)
	adapters.Register(adapters.Provider(aihub.Provider360), adapters.New360Adapter)
	adapters.Register(adapters.Provider(aihub.ProviderHunyuan), adapters.NewHunyuanAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderMoonshot), adapters.NewMoonshotAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderBaichuan), adapters.NewBaichuanAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderMiniMax), adapters.NewMinimaxAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderGroq), adapters.NewGroqAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderOllama), adapters.NewOllamaAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderYi), adapters.NewYiAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderStepFun), adapters.NewStepfunAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderCoze), adapters.NewCozeAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderCohere), adapters.NewCohereAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderTogether), adapters.NewTogetherAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderNovita), adapters.NewNovitaAdapter)
	adapters.Register(adapters.Provider(aihub.ProviderXAI), adapters.NewXaiAdapter)
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
