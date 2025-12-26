package lang

import (
	"math/rand"
)

// random byte array
func RandByteArray(length int) []int8 {
	datas := make([]int8, length)
	for i := 0; i < length; i++ {
		datas[i] = int8(rand.Intn(127) - 128)
	}
	return datas
}

// [min, max]  // 包含 min 和 max
func RandInt(min, max int) int {
	if min > max {
		panic("min cannot be greater than max")
	}
	return rand.Intn(max-min+1) + min
}
