package llmhub

import (
	"context"
	"io"

	"github.com/gotoailab/llmhub/internal/adapters"
	"github.com/gotoailab/llmhub/internal/models"
)

// adapterWrapper 包装内部适配器，将客户端类型转换为适配器类型
type adapterWrapper struct {
	adapter adapters.Adapter
}

func (w *adapterWrapper) ChatCompletion(ctx context.Context, req *internalChatCompletionRequest) (*internalChatCompletionResponse, error) {
	// 转换为适配器需要的类型
	adapterReq := &models.ChatCompletionRequest{
		Model:            req.Model,
		Messages:         w.toAdapterMessages(req.Messages),
		Temperature:      req.Temperature,
		TopP:             req.TopP,
		MaxTokens:        req.MaxTokens,
		Stream:           req.Stream,
		PresencePenalty:  req.PresencePenalty,
		FrequencyPenalty: req.FrequencyPenalty,
		Stop:             req.Stop,
		User:             req.User,
		Functions:        w.toAdapterFunctions(req.Functions),
		FunctionCall:     req.FunctionCall,
		LogitBias:        req.LogitBias,
		LogProbs:         req.LogProbs,
		TopLogProbs:      req.TopLogProbs,
		ResponseFormat:   w.toAdapterResponseFormat(req.ResponseFormat),
		Seed:             req.Seed,
		Tools:            w.toAdapterTools(req.Tools),
		ToolChoice:       req.ToolChoice,
	}

	resp, err := w.adapter.ChatCompletion(ctx, adapterReq)
	if err != nil {
		return nil, err
	}

	return w.toInternalResponse(resp), nil
}

func (w *adapterWrapper) ChatCompletionStream(ctx context.Context, req *internalChatCompletionRequest) (io.ReadCloser, error) {
	adapterReq := &models.ChatCompletionRequest{
		Model:            req.Model,
		Messages:         w.toAdapterMessages(req.Messages),
		Temperature:      req.Temperature,
		TopP:             req.TopP,
		MaxTokens:        req.MaxTokens,
		Stream:           true,
		PresencePenalty:  req.PresencePenalty,
		FrequencyPenalty: req.FrequencyPenalty,
		Stop:             req.Stop,
		User:             req.User,
		Functions:        w.toAdapterFunctions(req.Functions),
		FunctionCall:     req.FunctionCall,
		LogitBias:        req.LogitBias,
		LogProbs:         req.LogProbs,
		TopLogProbs:      req.TopLogProbs,
		ResponseFormat:   w.toAdapterResponseFormat(req.ResponseFormat),
		Seed:             req.Seed,
		Tools:            w.toAdapterTools(req.Tools),
		ToolChoice:       req.ToolChoice,
	}

	return w.adapter.ChatCompletionStream(ctx, adapterReq)
}

func (w *adapterWrapper) GetProvider() Provider {
	// 将 adapters.Provider 转换为 llmhub.Provider
	return Provider(w.adapter.GetProvider())
}

// 类型转换方法
func (w *adapterWrapper) toAdapterMessages(msgs []internalChatMessage) []models.ChatMessage {
	result := make([]models.ChatMessage, 0, len(msgs))
	for _, msg := range msgs {
		result = append(result, models.ChatMessage{
			Role:         msg.Role,
			Content:      msg.Content,
			Name:         msg.Name,
			FunctionCall: w.toAdapterFunctionCall(msg.FunctionCall),
			ToolCalls:    w.toAdapterToolCalls(msg.ToolCalls),
			ToolCallID:   msg.ToolCallID,
		})
	}
	return result
}

func (w *adapterWrapper) toAdapterFunctionCall(fc *internalFunctionCall) *models.FunctionCall {
	if fc == nil {
		return nil
	}
	return &models.FunctionCall{
		Name:      fc.Name,
		Arguments: fc.Arguments,
	}
}

func (w *adapterWrapper) toAdapterToolCalls(tcs []internalToolCall) []models.ToolCall {
	result := make([]models.ToolCall, 0, len(tcs))
	for _, tc := range tcs {
		result = append(result, models.ToolCall{
			ID:       tc.ID,
			Type:     tc.Type,
			Function: *w.toAdapterFunctionCall(&tc.Function),
		})
	}
	return result
}

func (w *adapterWrapper) toAdapterFunctions(funcs []internalFunctionDefinition) []models.FunctionDefinition {
	result := make([]models.FunctionDefinition, 0, len(funcs))
	for _, f := range funcs {
		result = append(result, models.FunctionDefinition{
			Name:        f.Name,
			Description: f.Description,
			Parameters:  f.Parameters,
		})
	}
	return result
}

func (w *adapterWrapper) toAdapterResponseFormat(rf *internalResponseFormat) *models.ResponseFormat {
	if rf == nil {
		return nil
	}
	return &models.ResponseFormat{
		Type: rf.Type,
	}
}

func (w *adapterWrapper) toAdapterTools(tools []internalTool) []models.Tool {
	result := make([]models.Tool, 0, len(tools))
	for _, tool := range tools {
		result = append(result, models.Tool{
			Type:     tool.Type,
			Function: w.toAdapterFunctions([]internalFunctionDefinition{tool.Function})[0],
		})
	}
	return result
}

func (w *adapterWrapper) toInternalResponse(resp *models.ChatCompletionResponse) *internalChatCompletionResponse {
	choices := make([]internalChatCompletionChoice, 0, len(resp.Choices))
	for _, choice := range resp.Choices {
		choices = append(choices, internalChatCompletionChoice{
			Index:        choice.Index,
			Message:      w.toInternalMessage(choice.Message),
			FinishReason: choice.FinishReason,
			Delta:        w.toInternalMessagePtr(choice.Delta),
		})
	}

	return &internalChatCompletionResponse{
		ID:                resp.ID,
		Object:            resp.Object,
		Created:           resp.Created,
		Model:             resp.Model,
		Choices:           choices,
		Usage:             w.toInternalUsage(resp.Usage),
		SystemFingerprint: resp.SystemFingerprint,
	}
}

func (w *adapterWrapper) toInternalMessage(msg models.ChatMessage) internalChatMessage {
	return internalChatMessage{
		Role:         msg.Role,
		Content:      msg.Content,
		Name:         msg.Name,
		FunctionCall: w.toInternalFunctionCall(msg.FunctionCall),
		ToolCalls:    w.toInternalToolCalls(msg.ToolCalls),
		ToolCallID:   msg.ToolCallID,
	}
}

func (w *adapterWrapper) toInternalMessagePtr(msg *models.ChatMessage) *internalChatMessage {
	if msg == nil {
		return nil
	}
	m := w.toInternalMessage(*msg)
	return &m
}

func (w *adapterWrapper) toInternalFunctionCall(fc *models.FunctionCall) *internalFunctionCall {
	if fc == nil {
		return nil
	}
	return &internalFunctionCall{
		Name:      fc.Name,
		Arguments: fc.Arguments,
	}
}

func (w *adapterWrapper) toInternalToolCalls(tcs []models.ToolCall) []internalToolCall {
	result := make([]internalToolCall, 0, len(tcs))
	for _, tc := range tcs {
		result = append(result, internalToolCall{
			ID:       tc.ID,
			Type:     tc.Type,
			Function: *w.toInternalFunctionCall(&tc.Function),
		})
	}
	return result
}

func (w *adapterWrapper) toInternalUsage(usage models.Usage) internalUsage {
	return internalUsage{
		PromptTokens:     usage.PromptTokens,
		CompletionTokens: usage.CompletionTokens,
		TotalTokens:      usage.TotalTokens,
	}
}
