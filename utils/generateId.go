package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateId() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
