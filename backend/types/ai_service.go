package types

// AICompletionResponse 定义了 AI 完成响应的结构
type AICompletionResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

// AIService 定义了 AI 服务的接口
type AIService interface {
	// OptimizeTextStream 流式优化文本
	OptimizeTextStream(text string, style TextStyle, sourceLanguage Language, targetLanguage Language, handler func(response AICompletionResponse)) (string, error)
}
