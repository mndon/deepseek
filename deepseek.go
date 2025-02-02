package deepseek

import (
	"context"

	"github.com/go-deepseek/deepseek/client"
	"github.com/go-deepseek/deepseek/config"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

const DEFAULT_TIMEOUT_SECONDS = 60

const (
	DEEPSEEK_CHAT_MODEL     = "deepseek-chat"
	DEEPSEEK_REASONER_MODEL = "deepseek-reasoner"
)

type Client interface {
	CallChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)
	CallChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)

	StreamChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)
	StreamChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)

	// PingChatCompletionsChat() (*DeepseekChatResponse, error) // TODO: VN -- impl
}

func NewClient(apiKey string) (Client, error) {
	config := config.Config{
		ApiKey:         apiKey,
		TimeoutSeconds: DEFAULT_TIMEOUT_SECONDS,
	}
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config config.Config) (Client, error) {
	return client.NewClient(config)
}
