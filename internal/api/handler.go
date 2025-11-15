package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aihub/internal/adapters"
	"github.com/aihub/internal/config"
	"github.com/aihub/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// ChatCompletions 处理聊天完成请求
func (h *Handler) ChatCompletions(c *gin.Context) {
	var req models.ChatCompletionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{
				Message: fmt.Sprintf("Invalid request: %v", err),
				Type:    "invalid_request_error",
			},
		})
		return
	}

	// 获取模型配置
	modelConfig := config.GetModelConfig(req.Model)
	if modelConfig == nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.ErrorDetail{
				Message: fmt.Sprintf("Model '%s' not found", req.Model),
				Type:    "invalid_request_error",
				Code:    "model_not_found",
			},
		})
		return
	}

	// 创建适配器（将字符串转换为 adapters.Provider）
	adapter, err := adapters.CreateAdapter(adapters.Provider(modelConfig.Provider), modelConfig.APIKey, modelConfig.BaseURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.ErrorDetail{
				Message: fmt.Sprintf("Failed to create adapter: %v", err),
				Type:    "internal_error",
			},
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	// 处理流式请求
	if req.Stream {
		h.handleStreamResponse(c, ctx, adapter, &req)
		return
	}

	// 处理非流式请求
	resp, err := adapter.ChatCompletion(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.ErrorDetail{
				Message: fmt.Sprintf("API error: %v", err),
				Type:    "api_error",
			},
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// handleStreamResponse 处理流式响应
func (h *Handler) handleStreamResponse(c *gin.Context, ctx context.Context, adapter adapters.Adapter, req *models.ChatCompletionRequest) {
	stream, err := adapter.ChatCompletionStream(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.ErrorDetail{
				Message: fmt.Sprintf("Stream error: %v", err),
				Type:    "api_error",
			},
		})
		return
	}
	defer stream.Close()

	// 设置 SSE 响应头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")

	c.Stream(func(w io.Writer) bool {
		buf := make([]byte, 4096)
		n, err := stream.Read(buf)
		if err != nil && err != io.EOF {
			return false
		}
		if n > 0 {
			// 发送 SSE 格式的数据
			fmt.Fprintf(w, "data: %s\n\n", string(buf[:n]))
			c.Writer.Flush()
		}
		return err == nil
	})
}

// Models 返回可用的模型列表
func (h *Handler) Models(c *gin.Context) {
	modelsList := make([]map[string]interface{}, 0, len(config.GlobalConfig.Models))

	for _, model := range config.GlobalConfig.Models {
		modelsList = append(modelsList, map[string]interface{}{
			"id":       model.Name,
			"object":   "model",
			"created":  time.Now().Unix(),
			"owned_by": model.Provider,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"object": "list",
		"data":   modelsList,
	})
}

// Health 健康检查
func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Unix(),
	})
}
