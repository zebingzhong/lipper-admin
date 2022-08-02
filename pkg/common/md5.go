package common

import (
	"crypto/md5"
	"encoding/hex"
)

//@autor: lipper
//@autor: MD5V
//@description: md5加密
//@param: str []byte
//@return: string
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
