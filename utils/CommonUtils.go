package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// md5加密
func MD5(str string) string {
	// md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	// return md5str
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 时间戳转字符串
func TimestampToStr(ts int64, str string) string {
	tm := time.Unix(ts, 0)
	return tm.Format(str)
}

// 字符串转时间戳
func StrToTimeStamp(str string, layout string) int64 {
	//获取本地location
	timeLocal, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(layout, str, timeLocal)
	return theTime.Unix()
}
