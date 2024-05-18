package v2

import "encoding/base64"

// Base64解码
func base64Decoding(str string) string {
	res, _ := base64.StdEncoding.DecodeString(str)
	return string(res)
}
