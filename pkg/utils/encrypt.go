package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptPassword(password string) string {
	return hashSHA256(password)
}

func hashSHA256(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	sha := hasher.Sum(nil)
	return hex.EncodeToString(sha)
}
