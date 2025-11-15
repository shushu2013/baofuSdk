package tool

import (
	"time"
)

func FormatDateTime(time time.Time, needSecond bool) string {
	if needSecond {
		return time.Format("2006-01-02 15:04:05")
	}
	return time.Format("2006-01-02 15:04")
}
