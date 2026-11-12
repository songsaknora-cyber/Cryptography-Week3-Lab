package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashSHA512(input string) string {
	hash := sha512.Sum512([]byte(input))
	return hex.EncodeToString(hash[:])
}