package pkg

import (
	"crypto/sha256"
	"encoding/hex"
)

func NewSHA256(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	sha256 := hash.Sum(nil)
	sha256String := hex.EncodeToString(sha256)
	return sha256String
}
