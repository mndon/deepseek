package deepseek

type ChatCompletionsResponse struct {
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
	Content          string `json:"content"`           // TODO: VN -- make it []byte
	ReasoningContent string `json:"reasoning_content"` // TODO: VN -- make it []byte
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
