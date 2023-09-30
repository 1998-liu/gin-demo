package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func MD5(str string) string {
	// md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	// return md5str
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
