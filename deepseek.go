package deepseek

import (
	"net/http"
	"time"

	"github.com/go-deepseek/deepseek/internal"
)

type Client interface {
	CallChatCompletionsChat(chatReq *ChatCompletionsRequest) (*ChatCompletionsResponse, error)
	CallChatCompletionsReasoner(chatReq *ChatCompletionsRequest) (*ChatCompletionsResponse, error)

	StreamChatCompletionsChat(chatReq *ChatCompletionsRequest) (*MessageIterator, error)
	StreamChatCompletionsReasoner(chatReq *ChatCompletionsRequest) (*MessageIterator, error)

	// PingChatCompletionsChat() (*DeepseekChatResponse, error) // TODO: VN -- impl
}

func NewClient(apiKey string) Client {
	return NewClientWithTimeout(apiKey, internal.DEFAULT_TIMEOUT_SECONDS)
}

func NewClientWithTimeout(apiKey string, timeoutSeconds int) Client {
	timeout := time.Second * time.Duration(timeoutSeconds)
	c := &client{
		ApiKey: apiKey,
		Client: &http.Client{
			Timeout: timeout,
		},
	}
	return c
}
