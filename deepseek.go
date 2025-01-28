package deepseek

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
	TopLogprobs int  `json:"top_logprobs,omitempty"`
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
	Logprobs     *Logprobs `json:"logprobs"`
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
