package util

import (
	"fmt"

	"github.com/brianvoe/sjwt"
)

func GenerateJwtToken(key string, value interface{}, secret string) string {
	claims := sjwt.New()
	claims.Set(key, value)
	secretKey := []byte(secret)

	return claims.Generate(secretKey)
}

func ParseJwtToken(key string, token string, secret string) (string, error) {
	secretKey := []byte(secret)

	if !sjwt.Verify(token, secretKey) {
		return "", fmt.Errorf("error to verify")
	}

	claims, err := sjwt.Parse(token)
	if err != nil {
		return "", fmt.Errorf("error to parse")
	}

	value, err := claims.GetStr(key)
	if err != nil {
		return "", fmt.Errorf("error to get")
	}

	return value, nil
}
