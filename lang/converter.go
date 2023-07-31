package lang

import "time"

type KnowedStructType interface {
	int8 | int16 | int32 | int64 | float32 | float64 | uint8 | uint16 | uint32 | uint64 |
		bool | time.Time
}

// 将bool值转换为指针
func BoolToPtr(v bool) *bool {
	return &v
}

func StringToPtr(v string) *string {
	return &v
}

// 将bool指针转换为bool,如果指针为nil，则直接返回false
func PtrToBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// 将int值转换为指针
func IntToPtr(v int) *int {
	return &v
}

// 将int指针转换为int,如果指针为nil，则直接返回0
func PtrToInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// 获取值的指针
func ValueToPtr[T any](v T) *T {
	return &v
}
