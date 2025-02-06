package infrastructure

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// ClaudeService handles communication with Claude API
type ClaudeService struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
	debug      bool
}

// ClaudeMessage represents a message in the chat
type ClaudeMessage struct {
	// 在Anthropic的API文档中，message结构中的role字段用于指定消息的角色。这个字段有两个可用的选项：user和assistant。以下是如何使用这些角色的说明：
	// 1. user:
	// 这个角色用于表示用户发送的消息。在对话中，用户的输入通常是问题、请求或指令。
	// 示例：
	// }
	// assistant:
	// 这个角色用于表示由Claude（或其他AI助手）生成的响应。在对话中，助手的输出通常是对用户输入的回答或执行的操作。
	// 示例：
	// }
	// 使用场景
	// 在构建对话时，role字段帮助区分对话的不同参与者。通常，一个完整的对话会交替包含user和assistant角色的消息。
	// 在发送请求时，你可以提供一系列的message对象，描述之前的对话历史。Claude会根据这些历史消息生成下一个响应。
	// 如果最后一条消息是assistant角色，Claude的响应将直接从该消息的内容继续。
	// 通过正确使用role字段，你可以有效地管理和构建与Claude的对话流。
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ClaudeCompletionRequest represents the request structure for chat completions
type ClaudeCompletionRequest struct {
	Model       string          `json:"model"`
	Messages    []ClaudeMessage `json:"messages"`
	Stream      bool            `json:"stream,omitempty"`
	Temperature float32         `json:"temperature,omitempty"`
	MaxTokens   int             `json:"max_tokens,omitempty"`
}

// ClaudeCompletionResponse represents the response structure from Claude
type ClaudeCompletionResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Created int64  `json:"created"`
	Object  string `json:"object"`
	Choices []struct {
		Index        string        `json:"index"`
		Delta        ClaudeMessage `json:"delta"`
		Message      ClaudeMessage `json:"message,omitempty"`
		FinishReason string        `json:"finish_reason"`
	} `json:"choices"`
	Usage *struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage,omitempty"`
}

// NewClaudeService creates a new Claude service instance
func NewClaudeService(authToken string) *ClaudeService {
	return &ClaudeService{
		baseURL:   "https://brconnector.runix.ai",
		authToken: authToken,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
		debug: true,
	}
}

func (s *ClaudeService) debugLog(format string, v ...interface{}) {
	if s.debug {
		log.Printf("[Claude Debug] "+format, v...)
	}
}

// CreateChatCompletion sends a chat completion request to Claude
func (s *ClaudeService) CreateChatCompletion(messages []ClaudeMessage, stream bool) (*ClaudeCompletionResponse, error) {
	url := fmt.Sprintf("%s/v1/chat/completions", s.baseURL)

	reqBody := ClaudeCompletionRequest{
		Model:    "claude-3-5-sonnet",
		Messages: messages,
		Stream:   stream,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize request: %w", err)
	}

	s.debugLog("Request body: %s", string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.authToken))

	s.debugLog("Sending request to: %s", url)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	s.debugLog("Response status: %s", resp.Status)
	for k, v := range resp.Header {
		s.debugLog("Response header %s: %v", k, v)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call failed, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	if stream {
		return nil, fmt.Errorf("streaming is not supported in this method, please use CreateStreamChatCompletion")
	}

	var result ClaudeCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	s.debugLog("Response body: %+v", result)
	return &result, nil
}

// CreateStreamChatCompletion sends a streaming chat completion request to Claude
func (s *ClaudeService) CreateStreamChatCompletion(messages []ClaudeMessage, handler func(response ClaudeCompletionResponse)) error {
	url := fmt.Sprintf("%s/v1/chat/completions", s.baseURL)

	reqBody := ClaudeCompletionRequest{
		Model:    "claude-3-5-sonnet",
		Messages: messages,
		Stream:   true,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to serialize request: %w", err)
	}

	s.debugLog("Stream request body: %s", string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.authToken))
	req.Header.Set("Accept", "text/event-stream")

	s.debugLog("Sending stream request to: %s", url)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	s.debugLog("Stream response status: %s", resp.Status)
	for k, v := range resp.Header {
		s.debugLog("Stream response header %s: %v", k, v)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API call failed, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		s.debugLog("Raw stream line: %q", line)

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "data:") {
			continue
		}

		data := strings.TrimPrefix(line, "data:")
		data = strings.TrimSpace(data)

		if data == "[DONE]" {
			s.debugLog("Received [DONE] message")
			break
		}

		var response ClaudeCompletionResponse
		if err := json.Unmarshal([]byte(data), &response); err != nil {
			s.debugLog("Failed to parse response data: %v, data: %s", err, data)
			continue
		}

		handler(response)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading stream: %w", err)
	}

	return nil
}
