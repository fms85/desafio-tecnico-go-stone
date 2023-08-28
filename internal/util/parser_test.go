package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHash(t *testing.T) {
	input := "hello"
	expectedHash := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"

	hash := GenerateHash(input)

	assert.Equal(t, expectedHash, hash, "Generated hash should match the expected hash")
}

func TestStringToUint(t *testing.T) {
	t.Run("should_convert_string_to_uint", func(t *testing.T) {
		input := "123"
		expectedOutput := uint(123)

		output := StringToUint(input)

		assert.Equal(t, expectedOutput, output, "Converted uint should match the expected value")
	})

	t.Run("should_panic_on_invalid_string", func(t *testing.T) {
		input := "invalid"

		assert.Panics(t, func() { StringToUint(input) }, "Invalid string should cause a panic")
	})
}
