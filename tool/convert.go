package tool

import (
	"strconv"
)

func Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}
