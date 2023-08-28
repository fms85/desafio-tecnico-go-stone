package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func GenerateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	hashString := hex.EncodeToString(hash[:])

	return hashString
}

func StringToUint(s string) uint {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return uint(i)
}
