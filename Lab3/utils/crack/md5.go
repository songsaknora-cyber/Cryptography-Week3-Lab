package crack

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}