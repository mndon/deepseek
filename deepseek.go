package deepseek

import (
	"errors"
	"fmt"
)

const (
	BASE_URL = `https://api.deepseek.com`
)

type DeepseekChatRequest struct {
	Messages         []*Message      `json:"messages"`
	Model            string          `json:"model"`
	FrequencyPenalty int             `json:"frequency_penalty,omitempty"`
	MaxTokens        int             `json:"max_tokens,omitempty"`
	PresencePenalty  int             `json:"presence_penalty,omitempty"`
	ResponseFormat   *ResponseFormat `json:"response_format,omitempty"`
	// stop // TODO: VN -- add support
	Stream        bool           `json:"stream,omitempty"`
	StreamOptions *StreamOptions `json:"stream_options,omitempty"`
	Temperature   int            `json:"temperature,omitempty"`
	TopP          int            `json:"top_p,omitempty"`
	// tools // TODO: VN -- add support
	// tool_choice // TODO: VN -- add support
	Logprobs    bool `json:"logprobs,omitempty"`
	TopLogprobs *int `json:"top_logprobs,omitempty"`
}

type Message struct {
	Role    string `json:"role"`    // TODO: VN -- support roles like system, user, assistant, tool
	Content string `json:"content"` // TODO: VN -- make it []byte
}

type ResponseFormat struct {
	Type string `json:"type"` // Must be one of text or json_object
}

type StreamOptions struct {
	IncludeUsage bool `json:"include_usage"`
}

type DeepseekChatResponse struct {
	Id                string    `json:"id"`
	Choices           []*Choice `json:"choices"`
	Created           int       `json:"created"`
	Model             string    `json:"model"`
	SystemFingerprint string    `json:"system_fingerprint"`
	Object            string    `json:"object"`
	Usage             *Usage    `json:"usage"`
}

type Choice struct {
	FinishReason string    `json:"finish_reason"`
	Index        int       `json:"index"`
	Message      *Message  `json:"message"`
	Delta        *Delta    `json:"delta"` // TODO: VN -- mesage and delta in one struct or diff?
	Logprobs     *Logprobs `json:"logprobs"`
}

type Delta struct {
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens          int                  `json:"prompt_tokens"`
	CompletionTokens      int                  `json:"completion_tokens"`
	TotalTokens           int                  `json:"total_tokens"`
	PromptTokensDetails   *PromptTokensDetails `json:"prompt_tokens_details"`
	PromptCacheHitTokens  int                  `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens int                  `json:"prompt_cache_miss_tokens"`
}

type PromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type Logprobs struct {
	Content *Content `json:"content"`
}

type Content struct {
	TopLogprob
	TopLogprobs []*TopLogprob `json:"top_logprobs"`
}

type TopLogprob struct {
	Token   string `json:"token"`
	Logprob int    `json:"logprob"`
	Bytes   []int  `json:"bytes"`
}

func ValidateRequest(req *DeepseekChatRequest) error {
	if req == nil {
		return errors.New("err: input request is nil")
	}
	if req.Messages == nil || len(req.Messages) < 1 {
		return errors.New("err: messages required in request")
	}
	if req.Model == "" {
		return errors.New("err: model required in request")
	}
	if !(req.FrequencyPenalty >= -2 && req.FrequencyPenalty <= 2) {
		return fmt.Errorf("err: frequency_penalty is invalid; it should be number between -2 and 2")
	}
	if !(req.MaxTokens == 0 || (req.MaxTokens >= 1 && req.MaxTokens <= 8192)) {
		return fmt.Errorf("err: max_tokens is invalid; it should be number between 1 and 8192 or 0")
	}
	if !(req.PresencePenalty >= -2 && req.PresencePenalty <= 2) {
		return fmt.Errorf("err: presence_penalty is invalid; it should be number between -2 and 2")
	}
	if req.ResponseFormat != nil {
		if !(req.ResponseFormat.Type == "text" || req.ResponseFormat.Type == "json_object") {
			return fmt.Errorf(`err: response_format type %q is invalid; it should be one of "text" or "json_object"`, req.ResponseFormat.Type)
		}
	}
	// TODO: VN -- stop validation
	if !req.Stream && req.StreamOptions != nil {
		return errors.New(`err: "stream_options" can not be set when "stream" is false`)
	}
	if !(req.Temperature >= 0 && req.Temperature <= 2) {
		return fmt.Errorf("err: temperature is invalid; it should be number between 0 and 2")
	}
	if !(req.TopP <= 1) {
		return fmt.Errorf("err: top_p is invalid; it should be number less than 1")
	}
	// TODO: VN -- tools validation
	// TODO: VN -- tool_choice validation
	if req.TopLogprobs != nil {
		if !req.Logprobs {
			return fmt.Errorf(`err: top_logprobs can not be set when "logprobs" is false`)
		}
		if !(*req.TopLogprobs >= 0 && *req.TopLogprobs <= 20) {
			return fmt.Errorf(`err: top_logprobs is invalid; it should be number between 0 and 20`)
		}
	}
	return nil
}
