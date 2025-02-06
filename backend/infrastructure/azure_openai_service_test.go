package infrastructure

import (
	"testing"
)

func TestAzureOpenAIService_CreateChatCompletion(t *testing.T) {
	service := NewAzureOpenAIService()

	messages := []ChatMessage{
		{
			Role:    "system",
			Content: "You are a helpful assistant.",
		},
		{
			Role:    "user",
			Content: "Hello! How are you?",
		},
	}

	response, err := service.CreateChatCompletion(messages)
	if err != nil {
		t.Errorf("CreateChatCompletion failed: %v", err)
		return
	}

	if len(response.Choices) == 0 {
		t.Error("Expected at least one choice in response")
		return
	}

	if response.Choices[0].Message.Content == "" {
		t.Error("Expected non-empty message content")
	}
}
