package hash

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

// Md5string returns an md5 has of the string passed in
func Md5strings(in string) string {
	h := md5.New()
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5r returns an md5 has of the interfaces passed in
func Md5r(ins ...interface{}) string {
	h := md5.New()
	for _, v := range ins {
		json.NewEncoder(h).Encode(v)
	}
	return hex.EncodeToString(h.Sum(nil))
}
