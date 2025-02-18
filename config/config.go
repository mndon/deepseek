package config

// Config for deepseek client.
//
//	ApiKey - deepseek API key.
//	TimeoutSeconds - http client timeout used by deepseek client.
//	DisableRequestValidation - disable request validation by deepseek client.
type Config struct {
	BaseURL                  string
	ApiKey                   string
	TimeoutSeconds           int
	DisableRequestValidation bool
}
