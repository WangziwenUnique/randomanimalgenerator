package infrastructure

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	// httpClient is a shared HTTP client with timeout
	httpClient = &http.Client{
		// Timeout: 120 * time.Second,
	}
)

// OllamaService handles communication with Ollama API
type OllamaService struct {
	baseURL string
}

// OllamaRequest represents the request structure sent to Ollama
type OllamaRequest struct {
	Model    string                 `json:"model"`
	Prompt   string                 `json:"prompt"`
	Stream   bool                   `json:"stream,omitempty"`
	Options  map[string]interface{} `json:"options,omitempty"`
	Context  []int                  `json:"context,omitempty"`
	System   string                 `json:"system,omitempty"`
	Template string                 `json:"template,omitempty"`
	Format   string                 `json:"format,omitempty"`
}

// OllamaResponse represents the response structure received from Ollama
type OllamaResponse struct {
	Model           string `json:"model"`
	CreatedAt       string `json:"created_at"`
	Response        string `json:"response"`
	Done            bool   `json:"done"`
	Context         []int  `json:"context,omitempty"`
	TotalDuration   int64  `json:"total_duration,omitempty"`
	LoadDuration    int64  `json:"load_duration,omitempty"`
	PromptEvalCount int    `json:"prompt_eval_count,omitempty"`
	EvalCount       int    `json:"eval_count,omitempty"`
	EvalDuration    int64  `json:"eval_duration,omitempty"`
}

// ResponseHandler handles the callback function for stream responses
type ResponseHandler func(response OllamaResponse)

// NewOllamaService creates a new Ollama service instance
func NewOllamaService() *OllamaService {
	// Get base URL from environment variable, fallback to default values
	env := os.Getenv("GO_ENV")
	baseURL := "http://localhost:11434"
	if env == "prod" {
		baseURL = "http://210.42.32.22:11455"
	}

	return &OllamaService{
		baseURL: baseURL,
	}
}

// IsTimeoutError checks if the error is a timeout error
func IsTimeoutError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "context deadline exceeded") ||
		strings.Contains(err.Error(), "Client.Timeout")
}

// GenerateStream sends a generate request to Ollama (stream)
func (s *OllamaService) GenerateStream(model, prompt string, handler ResponseHandler) error {
	reqBody := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: true,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to serialize request: %w", err)
	}

	req, err := http.NewRequest("POST", s.baseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API call failed, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var response OllamaResponse
		if err := json.Unmarshal(scanner.Bytes(), &response); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
		handler(response)
	}

	return scanner.Err()
}
