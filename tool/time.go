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
