package tool

import (
	"time"
)

func GetTimeMilliseconds(time time.Time) int64 {
	if time.IsZero() {
		return 0
	}
	return time.UnixMilli()
}

// GetReqTime 获取请求时间，格式：yyyyMMddHHmmss
func GetReqTime() string {
	return time.Now().Format("20060102150405")
}
