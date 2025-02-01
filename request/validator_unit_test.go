package request

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMessages(t *testing.T) {
	t.Run("err for zero messages", func(t *testing.T) {
		msgs := []*Message{}
		err := validateMessages(msgs)
		assert.NotNil(t, err)
	})

	t.Run("err for blank role", func(t *testing.T) {
		msgs := []*Message{
			{
				Content: "Hello",
			},
		}
		err := validateMessages(msgs)
		assert.NotNil(t, err)
	})

	t.Run("err for invalid role", func(t *testing.T) {
		msgs := []*Message{
			{
				Role: "random",
			},
		}
		err := validateMessages(msgs)
		assert.NotNil(t, err)
	})

	t.Run("err for role with caps", func(t *testing.T) {
		msgs := []*Message{
			{
				Role: strings.ToUpper(RoleUser),
			},
		}
		err := validateMessages(msgs)
		assert.NotNil(t, err)
	})

	t.Run("err for blank content", func(t *testing.T) {
		msgs := []*Message{
			{
				Role:    RoleUser,
				Content: "",
			},
		}
		err := validateMessages(msgs)
		assert.NotNil(t, err)
	})

	t.Run("err for blank tool_call_id", func(t *testing.T) {
		msgs := []*Message{
			{
				Role:    RoleTool,
				Content: "Hello",
			},
		}
		err := validateMessages(msgs)
		assert.NotNil(t, err)
	})

	t.Run("no err for blank tool_call_id", func(t *testing.T) {
		msgs := []*Message{
			{
				Role:    RoleUser,
				Content: "Hello",
			},
		}
		err := validateMessages(msgs)
		assert.NoError(t, err)
	})

	t.Run("no err with valid messages", func(t *testing.T) {
		msgs := []*Message{
			{
				Role:    RoleSystem,
				Content: "Hello from system",
			},
			{
				Role:    RoleUser,
				Content: "Hello from user",
			},
			{
				Role:    RoleAssistant,
				Content: "Hello from assistant",
			},
			{
				Role:       RoleTool,
				Content:    "Hello from tool",
				ToolCallId: "dummy",
			},
		}
		err := validateMessages(msgs)
		assert.NoError(t, err)
	})

}

func TestValidateModel(t *testing.T) {
	t.Run("err for model blank", func(t *testing.T) {
		req := &ChatCompletionsRequest{}
		err := validateModel(req)
		assert.NotNil(t, err)
	})

	t.Run("err for invalid model", func(t *testing.T) {
		req := &ChatCompletionsRequest{
			Model: "random",
		}
		err := validateModel(req)
		assert.NotNil(t, err)
	})

	t.Run("no err for valid model", func(t *testing.T) {
		req := &ChatCompletionsRequest{
			Model: modelChat,
		}
		err := validateModel(req)
		assert.NoError(t, err)

		req = &ChatCompletionsRequest{
			Model: modelReasoner,
		}
		err = validateModel(req)
		assert.NoError(t, err)
	})

}

func TestValidateResponseFormat(t *testing.T) {
	t.Run("err for invalid response_format", func(t *testing.T) {
		req := &ChatCompletionsRequest{
			ResponseFormat: &ResponseFormat{
				Type: "random",
			},
		}
		err := validateResponseFormat(req)
		assert.NotNil(t, err)
	})

	t.Run("no err for valid response_format", func(t *testing.T) {
		req := &ChatCompletionsRequest{
			ResponseFormat: &ResponseFormat{
				Type: ResponseFormatText,
			},
		}
		err := validateResponseFormat(req)
		assert.NoError(t, err)
	})
}

func TestValidateStreamOptions(t *testing.T) {
	t.Run("err for stream_options with stream is false", func(t *testing.T) {
		req := &ChatCompletionsRequest{
			StreamOptions: &StreamOptions{
				IncludeUsage: true,
			},
		}
		err := validateStreamOptions(req)
		assert.NotNil(t, err)
	})

	t.Run("no err for stream_options with stream is true", func(t *testing.T) {
		req := &ChatCompletionsRequest{
			Stream: true,
			StreamOptions: &StreamOptions{
				IncludeUsage: true,
			},
		}
		err := validateStreamOptions(req)
		assert.NoError(t, err)
	})
}
