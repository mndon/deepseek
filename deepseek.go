package deepseek

import (
	"net/http"
	"time"
)

type Client interface {
	CallChatCompletionsChat(chatReq *DeepseekChatRequest) (*DeepseekChatResponse, error)
	CallChatCompletionsReasoner(chatReq *DeepseekChatRequest) (*DeepseekChatResponse, error)

	StreamChatCompletionsChat(chatReq *DeepseekChatRequest) (*MessageIterator, error)
	StreamChatCompletionsReasoner(chatReq *DeepseekChatRequest) (*MessageIterator, error)
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
