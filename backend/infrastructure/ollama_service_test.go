package infrastructure

import (
	"fmt"
	"testing"
)

func TestOllamaService_GenerateStream(t *testing.T) {
	service := NewOllamaService()
	err := service.GenerateStream("llama3.2:3b", "test prompt", func(resp OllamaResponse) {
		fmt.Println(resp.Response)
	})
	if err != nil {
		t.Fatalf("GenerateStream failed: %v", err)
	}

}
