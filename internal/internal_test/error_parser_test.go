package internal_test

import (
	"testing"

	"github.com/go-deepseek/deepseek/internal"
	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {

	t.Run("no err for valid error response", func(t *testing.T) {
		body := `{
			"error": {
			  "message": "Authentication Fails (no such user)",
			  "type": "authentication_error",
			  "param": null,
			  "code": "invalid_request_error"
			}
		  }`

		got, err := internal.ParseError([]byte(body))
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, `Authentication Fails (no such user)`, got.Error.Message)
		assert.Equal(t, `authentication_error`, got.Error.Type)
		assert.Equal(t, `invalid_request_error`, got.Error.Code)
	})

	t.Run("err for invalid non-json error response", func(t *testing.T) {
		body := `Invalid request`
		_, err := internal.ParseError([]byte(body))
		assert.Error(t, err)
	})

}
