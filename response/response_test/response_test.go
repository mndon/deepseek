package response_test

import (
	"encoding/json"
	"testing"

	"github.com/go-deepseek/deepseek/response"
	"github.com/stretchr/testify/assert"
)

func TestChatResponse_LoadJson(t *testing.T) {
	content := `{
    "id": "7713fabe-5401-4fbb-905f-0bead65ce42f",
    "object": "chat.completion",
    "created": 1738119582,
    "model": "deepseek-chat",
    "choices": [
        {
            "index": 0,
            "message": {
                "role": "assistant",
                "content": "Hello! How can I assist you today? 😊"
            },
            "logprobs": null,
            "finish_reason": "stop"
        }
    ],
    "usage": {
        "prompt_tokens": 11,
        "completion_tokens": 11,
        "total_tokens": 22,
        "prompt_tokens_details": {
            "cached_tokens": 0
        },
        "prompt_cache_hit_tokens": 0,
        "prompt_cache_miss_tokens": 11
    },
    "system_fingerprint": "fp_3a5770e1b4"
}`
	chatResp := &response.ChatCompletionsResponse{}

	err := json.Unmarshal([]byte(content), chatResp)

	assert.NoError(t, err)

	// assert chatResp
	assert.NotNil(t, chatResp)
	assert.Equal(t, chatResp.Id, "7713fabe-5401-4fbb-905f-0bead65ce42f")
	assert.Equal(t, chatResp.Object, "chat.completion")
	assert.Equal(t, chatResp.Created, 1738119582)
	assert.Equal(t, "deepseek-chat", chatResp.Model)
	assert.Equal(t, "fp_3a5770e1b4", chatResp.SystemFingerprint)

	// assert choices
	assert.Len(t, chatResp.Choices, 1)
	choice := chatResp.Choices[0]
	assert.Equal(t, 0, choice.Index)
	assert.NotNil(t, choice.Message)
	assert.Equal(t, "assistant", choice.Message.Role)
	assert.Equal(t, "Hello! How can I assist you today? 😊", choice.Message.Content)
	assert.Equal(t, "stop", choice.FinishReason)

	// assert usage
	assert.NotNil(t, chatResp.Usage)
	assert.Equal(t, 11, chatResp.Usage.PromptTokens)
	assert.Equal(t, 11, chatResp.Usage.CompletionTokens)
	assert.Equal(t, 22, chatResp.Usage.TotalTokens)
	assert.NotNil(t, chatResp.Usage.PromptTokensDetails)
	assert.Equal(t, 0, chatResp.Usage.PromptTokensDetails.CachedTokens)
	assert.Equal(t, 0, chatResp.Usage.PromptCacheHitTokens)
	assert.Equal(t, 11, chatResp.Usage.PromptCacheMissTokens)
}
