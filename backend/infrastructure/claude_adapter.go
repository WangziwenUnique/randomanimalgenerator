package infrastructure

import (
	"fmt"
	"strings"

	"github.com/aogen-fiber/backend/types"
)

// ClaudeAdapter 将 ClaudeService 适配到 AIService 接口
type ClaudeAdapter struct {
	claude *ClaudeService
}

// NewClaudeAdapter 创建新的 Claude 适配器
func NewClaudeAdapter(claude *ClaudeService) *ClaudeAdapter {
	return &ClaudeAdapter{
		claude: claude,
	}
}

// OptimizeTextStream 实现带流式返回的文本优化
func (a *ClaudeAdapter) OptimizeTextStream(text string, style types.TextStyle, sourceLanguage types.Language, targetLanguage types.Language, handler func(response types.AICompletionResponse)) (string, error) {
	messages := []ClaudeMessage{
		{
			Role:    "user",
			Content: style.GetPrompt(text, sourceLanguage, targetLanguage),
		},
	}

	// 创建一个通道来接收完整的输出文本
	var fullOutput strings.Builder

	// 处理流式响应
	err := a.claude.CreateStreamChatCompletion(messages, func(response ClaudeCompletionResponse) {
		if len(response.Choices) > 0 {
			content := response.Choices[0].Delta.Content
			if content != "" {
				fullOutput.WriteString(content)
			}
			// 转换为通用的 AICompletionResponse
			handler(types.AICompletionResponse{
				ID:      response.ID,
				Content: content,
				Done:    response.Choices[0].FinishReason != "",
			})
		}
	})

	if err != nil {
		return "", fmt.Errorf("failed to optimize text: %w", err)
	}

	return fullOutput.String(), nil
}
