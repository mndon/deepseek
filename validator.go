package deepseek

import (
	"errors"
	"fmt"
)

func ValidateChatCompletionsRequest(req *ChatCompletionsRequest) error {
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
