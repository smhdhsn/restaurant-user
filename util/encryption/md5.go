package encryption

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 is responsible for encrypting a given string with md5 algorythm.
func EncodeMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
