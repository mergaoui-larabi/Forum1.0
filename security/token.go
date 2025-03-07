package security

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal("failed to generat token %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
