package deepseek

import (
	"net/http"
	"time"

	"github.com/go-deepseek/deepseek/client"
	"github.com/go-deepseek/deepseek/internal"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

type Client interface {
	CallChatCompletionsChat(chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)
	CallChatCompletionsReasoner(chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)

	StreamChatCompletionsChat(chatReq *request.ChatCompletionsRequest) (*response.MessageIterator, error)
	StreamChatCompletionsReasoner(chatReq *request.ChatCompletionsRequest) (*response.MessageIterator, error)

	// PingChatCompletionsChat() (*DeepseekChatResponse, error) // TODO: VN -- impl
}

func NewClient(apiKey string) Client {
	return NewClientWithTimeout(apiKey, internal.DEFAULT_TIMEOUT_SECONDS)
}

func NewClientWithTimeout(apiKey string, timeoutSeconds int) Client {
	timeout := time.Second * time.Duration(timeoutSeconds)
	c := &client.Client{
		ApiKey: apiKey,
		Client: &http.Client{
			Timeout: timeout,
		},
	}
	return c
}
