package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndParseJwtToken(t *testing.T) {
	key := "user_id"
	value := "123"
	secret := "mySecretKey"

	token := GenerateJwtToken(key, value, secret)

	t.Run("should_generate_and_parse_jwt_token", func(t *testing.T) {
		parsedValue, err := ParseJwtToken(key, token, secret)
		assert.NoError(t, err, "Parsing should not return an error")
		assert.Equal(t, value, parsedValue, "Parsed value should match the original value")
	})

	t.Run("should_fail_to_parse_invalid_token", func(t *testing.T) {
		invalidToken := "invalid_token"
		parsedValue, err := ParseJwtToken(key, invalidToken, secret)
		assert.Error(t, err, "Parsing invalid token should return an error")
		assert.Empty(t, parsedValue, "Parsed value from invalid token should be empty")
	})
}
