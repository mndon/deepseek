package deepseek

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
