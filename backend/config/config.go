package config

import (
	"os"
	"sync"
)

var (
	isProd bool
	mu     sync.RWMutex
)

// SetProd 设置生产环境标志
func SetProd(prod bool) {
	mu.Lock()
	defer mu.Unlock()
	isProd = prod
}

// IsProd 获取当前是否为生产环境
func IsProd() bool {
	mu.RLock()
	defer mu.RUnlock()
	return isProd
}

// Config 应用配置
type Config struct {
	GoogleClientID       string
	ClaudeAuthToken      string
	AzureOpenAIEndpoint  string
	AzureOpenAIAPIKey    string
	AzureOpenAIModel     string
	StripeSecretKey      string
	StripePublishableKey string
}

// LoadConfig 从环境变量加载配置
func LoadConfig() *Config {
	return &Config{
		GoogleClientID:       getEnvWithDefault("GOOGLE_CLIENT_ID", ""),
		ClaudeAuthToken:      getEnvWithDefault("CLAUDE_AUTH_TOKEN", ""),
		AzureOpenAIEndpoint:  getEnvWithDefault("AZURE_OPENAI_ENDPOINT", ""),
		AzureOpenAIAPIKey:    getEnvWithDefault("AZURE_OPENAI_API_KEY", ""),
		AzureOpenAIModel:     getEnvWithDefault("AZURE_OPENAI_MODEL", ""),
		StripeSecretKey:      getEnvWithDefault("STRIPE_SECRET_KEY", ""),
		StripePublishableKey: getEnvWithDefault("STRIPE_PUBLISHABLE_KEY", ""),
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
