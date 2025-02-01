package deepseek

import (
	"github.com/go-deepseek/deepseek/client"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

const DEFAULT_TIMEOUT_SECONDS = 60

const (
	DEEPSEEK_CHAT_MODEL     = "deepseek-chat"
	DEEPSEEK_REASONER_MODEL = "deepseek-reasoner"
)

type Client interface {
	CallChatCompletionsChat(chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)
	CallChatCompletionsReasoner(chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)

	StreamChatCompletionsChat(chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)
	StreamChatCompletionsReasoner(chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)

	// PingChatCompletionsChat() (*DeepseekChatResponse, error) // TODO: VN -- impl
}

func NewClient(apiKey string) Client {
	return NewClientWithTimeout(apiKey, DEFAULT_TIMEOUT_SECONDS)
}

func NewClientWithTimeout(apiKey string, timeoutSeconds int) Client {
	return client.NewClient(apiKey, timeoutSeconds)
}
