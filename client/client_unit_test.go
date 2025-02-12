package client

import (
	"testing"

	"github.com/go-deepseek/deepseek/request"
	"github.com/stretchr/testify/assert"
)

func TestValidateChatParams(t *testing.T) {

	t.Run("no err for valid req", func(t *testing.T) {
		chatReq := &request.ChatCompletionsRequest{
			Model:  "deepseek-chat",
			Stream: false,
		}
		err := validateChatParams(chatReq, false, "deepseek-chat")
		assert.NoError(t, err)
	})

	t.Run("err for invalid req having wrong model", func(t *testing.T) {
		chatReq := &request.ChatCompletionsRequest{
			Model:  "deepseek-reasoner",
			Stream: false,
		}
		err := validateChatParams(chatReq, false, "deepseek-chat")
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "deepseek-chat")
	})

	t.Run("err for invalid req having wrong stream", func(t *testing.T) {
		chatReq := &request.ChatCompletionsRequest{
			Model:  "deepseek-chat",
			Stream: true,
		}
		err := validateChatParams(chatReq, false, "deepseek-chat")
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "false")
	})

}
