package api

import (
	"github.com/aihub/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handler) *gin.Engine {
	router := gin.Default()

	// 健康检查（不需要认证）
	router.GET("/health", handler.Health)

	// API 路由组（需要认证）
	api := router.Group("/v1")
	api.Use(auth.APIKeyAuth())
	{
		api.POST("/chat/completions", handler.ChatCompletions)
		api.GET("/models", handler.Models)
	}

	return router
}

