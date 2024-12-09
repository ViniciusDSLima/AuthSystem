package utils

import (
	"encoding/hex"
	"math/rand"
)

func GenerateRandomToken(length int) string {
	bytes := make([]byte, length)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
