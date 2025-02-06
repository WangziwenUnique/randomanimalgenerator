package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aogen-fiber/backend/config"
)

// AzureOpenAIService handles communication with Azure OpenAI API
type AzureOpenAIService struct {
	endpoint   string
	apiKey     string
	model      string
	httpClient *http.Client
}

// ChatMessage represents a message in the chat
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionRequest represents the request structure for chat completions
type ChatCompletionRequest struct {
	Messages    []ChatMessage `json:"messages"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	Temperature float32       `json:"temperature,omitempty"`
}

// ChatCompletionResponse represents the response structure from Azure OpenAI
type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Message      ChatMessage `json:"message"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// NewAzureOpenAIService creates a new Azure OpenAI service instance
func NewAzureOpenAIService() *AzureOpenAIService {
	cfg := config.LoadConfig()
	return &AzureOpenAIService{
		endpoint: cfg.AzureOpenAIEndpoint,
		apiKey:   cfg.AzureOpenAIAPIKey,
		model:    cfg.AzureOpenAIModel,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// CreateChatCompletion sends a chat completion request to Azure OpenAI
func (s *AzureOpenAIService) CreateChatCompletion(messages []ChatMessage) (*ChatCompletionResponse, error) {
	url := fmt.Sprintf("%s/openai/deployments/%s/chat/completions?api-version=2024-05-01-preview", s.endpoint, s.model)

	reqBody := ChatCompletionRequest{
		Messages: messages,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", s.apiKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API call failed, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	var result ChatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CreateStreamChatCompletion sends a streaming chat completion request to Azure OpenAI
func (s *AzureOpenAIService) CreateStreamChatCompletion(messages []ChatMessage, handler func(response ChatCompletionResponse)) error {
	// TODO: Implement streaming chat completion
	return fmt.Errorf("streaming not implemented yet")
}
