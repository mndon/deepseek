package deepseek

import (
	"net/http"
	"time"
)

type Client interface {
	CallChatCompletionsChat(chatReq *DeepseekChatRequest) (*DeepseekChatResponse, error)
	// CallChatCompletionsReasoner()

	StreamChatCompletionsChat(chatReq *DeepseekChatRequest) (*MessageIterator, error)
	// StreamChatCompletionsReasoner()
}

func NewClient(apiKey string) Client {
	return NewClientWithTimeout(apiKey, DEFAULT_TIMEOUT_SECONDS)
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
