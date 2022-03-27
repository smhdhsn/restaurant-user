package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

// ToMD5 is responsible for encrypting a given string with md5 algorythm.
func ToMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
