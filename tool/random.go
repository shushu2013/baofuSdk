package tool

import (
	"fmt"
	"math"
	"math/rand"
)

func RandomInt(length int) int64 {
	min, max := int64(math.Pow10(length-1)), int64(math.Pow10(length)-1)
	return rand.Int63n(max-min) + min
}

// 结果是 [min, max)
func RandomRange(min, max int) int {
	if min == max {
		return min
	}
	value := min + rand.Intn(max-min)
	return value
}

func RandomStr(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	minIndex := 0
	maxIndex := len(chars) - 1
	result := ""
	for i := 0; i < length; i++ {
		result += fmt.Sprintf("%c", chars[RandomRange(minIndex, maxIndex)])
	}
	return result
}
