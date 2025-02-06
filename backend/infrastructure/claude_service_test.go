package infrastructure

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestClaudeService_CreateChatCompletion(t *testing.T) {
	t.Log("Starting TestClaudeService_CreateChatCompletion...")

	// 这里需要替换为实际的认证令牌
	authToken := "br-ykOFENNouxvfvrdFHrDRmz2oNac1p"
	service := NewClaudeService(authToken)
	t.Log("Created Claude service instance")

	messages := []ClaudeMessage{
		{
			Role:    "user",
			Content: "Hello, how are you?",
		},
	}

	// 打印请求信息
	requestJSON, _ := json.MarshalIndent(messages, "", "  ")
	t.Logf("Sending request with messages: %s", string(requestJSON))

	response, err := service.CreateChatCompletion(messages, false)
	if err != nil {
		t.Errorf("CreateChatCompletion failed: %v", err)
		return
	}

	// 打印完整的响应信息
	responseJSON, _ := json.MarshalIndent(response, "", "  ")
	t.Logf("Received response: %s", string(responseJSON))

	if len(response.Choices) == 0 {
		t.Error("Expected at least one choice in response")
		return
	}

	t.Logf("Response message content: %s", response.Choices[0].Message.Content)
	t.Logf("Model used: %s", response.Model)
	t.Logf("Finish reason: %s", response.Choices[0].FinishReason)

	if response.Choices[0].Message.Content == "" {
		t.Error("Expected non-empty message content")
	}

	t.Log("TestClaudeService_CreateChatCompletion completed successfully")
}

func TestClaudeService_CreateStreamChatCompletion(t *testing.T) {
	t.Log("Starting TestClaudeService_CreateStreamChatCompletion...")

	// 这里需要替换为实际的认证令牌
	authToken := "br-ykOFENNouxvfvrdFHrDRmz2oNac1p"
	service := NewClaudeService(authToken)
	t.Log("Created Claude service instance")

	testCases := []struct {
		name     string
		messages []ClaudeMessage
	}{
		{
			name: "Simple greeting",
			messages: []ClaudeMessage{
				{
					Role:    "user",
					Content: "Hello, how are you?",
				},
			},
		},
		{
			name: "Multiple messages",
			messages: []ClaudeMessage{
				{
					Role:    "system",
					Content: "You are a helpful assistant.",
				},
				{
					Role:    "user",
					Content: "What's the weather like?",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 打印请求信息
			reqBody := ClaudeCompletionRequest{
				Model:    "claude-3-5-sonnet",
				Messages: tc.messages,
				Stream:   true,
			}
			requestJSON, _ := json.MarshalIndent(reqBody, "", "  ")
			t.Logf("Sending streaming request: %s", string(requestJSON))

			messageCount := 0
			var receivedResponse bool
			var lastError error
			var fullContent strings.Builder
			done := make(chan bool)

			go func() {
				err := service.CreateStreamChatCompletion(tc.messages, func(response ClaudeCompletionResponse) {
					receivedResponse = true
					messageCount++

					// 打印每个流式响应
					responseJSON, _ := json.MarshalIndent(response, "", "  ")
					t.Logf("Received stream response #%d: %s", messageCount, string(responseJSON))

					if len(response.Choices) == 0 {
						lastError = fmt.Errorf("expected at least one choice in stream response")
						return
					}

					// 处理增量内容
					if response.Choices[0].Delta.Content != "" {
						fullContent.WriteString(response.Choices[0].Delta.Content)
						t.Logf("Stream message #%d delta content: %s", messageCount, response.Choices[0].Delta.Content)
					}

					if response.Choices[0].FinishReason != "" {
						t.Logf("Stream message #%d finish reason: %s", messageCount, response.Choices[0].FinishReason)
					}

					if response.Usage != nil {
						t.Logf("Stream message #%d usage - prompt: %d, completion: %d, total: %d",
							messageCount,
							response.Usage.PromptTokens,
							response.Usage.CompletionTokens,
							response.Usage.TotalTokens)
					}
				})
				if err != nil {
					lastError = err
				}
				done <- true
			}()

			// 设置超时
			select {
			case <-done:
				if lastError != nil {
					t.Errorf("CreateStreamChatCompletion failed: %v", lastError)
					return
				}
			case <-time.After(30 * time.Second):
				t.Error("Test timed out after 30 seconds")
				return
			}

			t.Logf("Total stream messages received: %d", messageCount)
			t.Logf("Complete response content: %s", fullContent.String())

			if !receivedResponse {
				t.Error("Expected to receive at least one response in stream")
			}

			if messageCount == 0 {
				t.Error("Expected to receive at least one message")
			}

			if fullContent.Len() == 0 {
				t.Error("Expected non-empty response content")
			}
		})
	}

	t.Log("TestClaudeService_CreateStreamChatCompletion completed successfully")
}

// 辅助函数：打印格式化的JSON
func prettyJSON(data interface{}) string {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v", err)
	}
	return string(bytes)
}
