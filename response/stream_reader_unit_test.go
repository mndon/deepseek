package response

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessResponse(t *testing.T) {
	t.Run("response keep-alive return error", func(t *testing.T) {
		respBody := []byte(KEEP_ALIVE)
		_, err := processResponse(respBody)
		assert.Error(t, err)
	})

	t.Run("response done return error", func(t *testing.T) {
		respBody := []byte(`data: [DONE]`)
		_, err := processResponse(respBody)
		assert.Error(t, err)
		assert.Equal(t, err, io.EOF)
	})

	t.Run("response json return chat response", func(t *testing.T) {
		respBody := []byte(`data: {"id":"aceb72f7-ffab-422a-b498-62c9b4034f84","object":"chat.completion.chunk","created":1738119601,"model":"deepseek-chat","system_fingerprint":"fp_3a5770e1b4","choices":[{"index":0,"delta":{"role":"assistant","content":""},"logprobs":null,"finish_reason":null}]}`)
		chatResp, err := processResponse(respBody)
		assert.NoError(t, err)
		assert.NotNil(t, chatResp)
		assert.Equal(t, "aceb72f7-ffab-422a-b498-62c9b4034f84", chatResp.Id)
	})
}

func TestTrimDataPrefix(t *testing.T) {
	t.Run("data prefix trimmed from json response", func(t *testing.T) {
		dataPrefix := `data: `
		jsonResp := `{"id":"aceb72f7-ffab-422a-b498-62c9b4034f84","object":"chat.completion.chunk","created":1738119601,"model":"deepseek-chat","system_fingerprint":"fp_3a5770e1b4","choices":[{"index":0,"delta":{"role":"assistant","content":""},"logprobs":null,"finish_reason":null}]}`
		respBody := []byte(dataPrefix + jsonResp)
		gotBody := trimDataPrefix(respBody)
		assert.Equal(t, jsonResp, string(gotBody))
	})

	t.Run("data prefix not trimmed from done response", func(t *testing.T) {
		dataPrefix := `data: `
		doneResp := `[DONE]`
		respBody := []byte(dataPrefix + doneResp)
		gotBody := trimDataPrefix(respBody)
		assert.Equal(t, doneResp, string(gotBody))
	})
}
