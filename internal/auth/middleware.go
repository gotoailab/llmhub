package auth

import (
	"net/http"
	"strings"

	"github.com/aihub/internal/config"
	"github.com/gin-gonic/gin"
)

// APIKeyAuth 中间件：验证 API Key
func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"message": "Missing Authorization header",
					"type":    "invalid_request_error",
					"code":    "missing_authorization",
				},
			})
			c.Abort()
			return
		}

		// 提取 Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"message": "Invalid Authorization header format",
					"type":    "invalid_request_error",
					"code":    "invalid_authorization",
				},
			})
			c.Abort()
			return
		}

		apiKey := parts[1]
		
		// 验证 API Key
		if !isValidAPIKey(apiKey) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{
					"message": "Invalid API key",
					"type":    "invalid_request_error",
					"code":    "invalid_api_key",
				},
			})
			c.Abort()
			return
		}

		// 将 API Key 存储到上下文中，供后续使用
		c.Set("api_key", apiKey)
		c.Next()
	}
}

// isValidAPIKey 验证 API Key 是否有效
func isValidAPIKey(apiKey string) bool {
	if config.GlobalConfig == nil {
		return false
	}

	// 检查配置中的 API Keys
	for _, validKey := range config.GlobalConfig.Auth.APIKeys {
		if apiKey == validKey {
			return true
		}
	}

	// 也可以检查模型配置中的 API Key（用于直接使用模型 API Key 的情况）
	for _, model := range config.GlobalConfig.Models {
		if apiKey == model.APIKey {
			return true
		}
	}

	return false
}

