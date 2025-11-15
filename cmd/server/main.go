package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aihub/internal/adapters"
	"github.com/aihub/internal/api"
	"github.com/aihub/internal/config"
)

func init() {
	// 注册所有适配器
	adapters.Register("openai", adapters.NewOpenAIAdapter)
	adapters.Register("claude", adapters.NewClaudeAdapter)
	adapters.Register("deepseek", adapters.NewDeepSeekAdapter)
	adapters.Register("qwen", adapters.NewQwenAdapter)
	adapters.Register("siliconflow", adapters.NewSiliconFlowAdapter)

	// 注册新适配器
	adapters.Register("gemini", adapters.NewGeminiAdapter)
	adapters.Register("mistral", adapters.NewMistralAdapter)
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
	// Cloudflare 和 DeepL 需要特殊处理，暂时跳过
	// adapters.Register("cloudflare", adapters.NewCloudflareAdapter)
	// adapters.Register("deepl", adapters.NewDeepLAdapter)
	adapters.Register("together", adapters.NewTogetherAdapter)
	adapters.Register("novita", adapters.NewNovitaAdapter)
	adapters.Register("xai", adapters.NewXaiAdapter)
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
