package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCPF(t *testing.T) {
	validCPF := "25462557035"
	invalidCPF := "43243243243"

	t.Run("should_return_true_for_valid_cpf", func(t *testing.T) {
		valid := ValidateCPF(validCPF)
		assert.True(t, valid, "Valid CPF should return true")
	})

	t.Run("should_return_false_for_invalid_cpf", func(t *testing.T) {
		invalid := ValidateCPF(invalidCPF)
		assert.False(t, invalid, "Invalid CPF should return false")
	})

	t.Run("should_return_false_for_invalid_length_cpf", func(t *testing.T) {
		invalidCPF := "1234567890" // Invalid length (10 characters)
		invalid := ValidateCPF(invalidCPF)
		assert.False(t, invalid, "Invalid length CPF should return false")
	})

	t.Run("should_return_zero_for_mismatched_length_input", func(t *testing.T) {
		invalidInput := "12345" // Mismatched length for the table
		invalidSum := validateCPFSumDigit(invalidInput, cpfFirstDigitTable)
		assert.Equal(t, 0, invalidSum, "Mismatched length input should return zero sum")
	})
}
