package tool

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func Base64DecodeStr(str string) (string, error) {
	ret, err := base64.StdEncoding.DecodeString(str)
	return string(ret), err
}
