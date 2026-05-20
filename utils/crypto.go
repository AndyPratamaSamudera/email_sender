package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashSHA512(input string) string {
	hasher := sha512.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ValidateAccessCode(providedCode string, validHash string) bool {
	hashedProvided := HashSHA512(providedCode)
	return hashedProvided == validHash
}
