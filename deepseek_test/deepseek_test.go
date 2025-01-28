package deepseek_test

import (
	"embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-deepseek/deepseek"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const TEST_API_KEY = `sk-123cd456c78d9be0b123de45cf6789b0` // replace with valid one

//go:embed testdata/*
var testdata embed.FS

func TestDeepseekChat(t *testing.T) {
	client := deepseek.NewClient(TEST_API_KEY)

	reqJson, err := testdata.ReadFile("testdata/01_req_basic_chat.json")
	assert.NoError(t, err)
	req := &deepseek.DeepseekChatRequest{}
	err = json.Unmarshal(reqJson, req)
	assert.NoError(t, err)

	resp, err := client.Call(req) // test

	require.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
}

func TestDeepseekChatStream(t *testing.T) {
	client := deepseek.NewClient(TEST_API_KEY)

	reqJson, err := testdata.ReadFile("testdata/02_req_stream_chat.json")
	assert.NoError(t, err)
	req := &deepseek.DeepseekChatRequest{}
	err = json.Unmarshal(reqJson, req)
	assert.NoError(t, err)

	iter, err := client.Stream(req) // test

	require.NoError(t, err)
	assert.NotNil(t, iter)

	for {
		resp := iter.Next()
		if resp == nil {
			break
		}
		fmt.Print(resp.Choices[0].Delta.Content)
	}
}

func TestResponse(t *testing.T) {
	respJson := `{
  "id": "dummy_string",
  "choices": [
    {
      "finish_reason": "stop",
      "index": 1,
      "message": {
        "content": "dummy_string",
        "reasoning_content": "dummy_string",
        "tool_calls": [
          {
            "id": "dummy_string",
            "type": "function",
            "function": {
              "name": "dummy_string",
              "arguments": "dummy_string"
            }
          }
        ],
        "role": "assistant"
      },
      "logprobs": {
        "content": [
          {
            "token": "dummy_string",
            "logprob": 1,
            "bytes": [
              1
            ],
            "top_logprobs": [
              {
                "token": "dummy_string",
                "logprob": 1,
                "bytes": [
                  1
                ]
              }
            ]
          }
        ]
      }
    }
  ],
  "created": 1,
  "model": "dummy_string",
  "system_fingerprint": "dummy_string",
  "object": "chat.completion",
  "usage": {
    "completion_tokens": 1,
    "prompt_tokens": 1,
    "prompt_cache_hit_tokens": 1,
    "prompt_cache_miss_tokens": 1,
    "total_tokens": 1,
    "completion_tokens_details": {
      "reasoning_tokens": 1
    }
  }
}`

	resp := &deepseek.DeepseekChatResponse{}
	json.Unmarshal([]byte(respJson), resp)

	wantStr := "dummy_string"
	wantInt := 1

	assert.NotNil(t, resp)

	assert.Equal(t, resp.Id, wantStr)
	assert.Len(t, resp.Choices, 1)
	assert.Equal(t, resp.Created, wantInt)
	assert.Equal(t, resp.Model, wantStr)
	assert.Equal(t, resp.SystemFingerprint, wantStr)
	assert.Equal(t, resp.Object, "chat.completion")
	assert.NotNil(t, resp.Usage)

	choice := resp.Choices[0]
	assert.Equal(t, choice.FinishReason, wantStr)
	assert.Equal(t, choice.Index, wantInt)
	assert.NotNil(t, choice.Message)

	message := choice.Message
	assert.Equal(t, message.Content, wantStr)
	assert.Equal(t, message.Role, wantStr)
	// TODO: VN -- complete reasoning_content

	usage := resp.Usage
	assert.Equal(t, usage.CompletionTokens, wantInt)
	assert.Equal(t, usage.PromptTokens, wantInt)
	assert.Equal(t, usage.PromptCacheHitTokens, wantInt)
	assert.Equal(t, usage.PromptCacheMissTokens, wantInt)
	assert.Equal(t, usage.TotalTokens, wantInt)
	assert.NotNil(t, usage.PromptTokensDetails)

	tokenDetails := usage.PromptTokensDetails
	assert.Equal(t, tokenDetails.CachedTokens, wantInt)
}
