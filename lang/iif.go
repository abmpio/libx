package lang

// 类似于c#中的? :赋值运算，返回一个指针
func IfValuePtr[K any](cond bool, trueValue func() *K, falseValue *K) *K {
	if cond {
		return trueValue()
	}
	return falseValue
}

// 类似于c#中的? :赋值运算，返回一个值
func IfValue[K any](cond bool, trueValueCallback func() K, falseValue K) K {
	if cond {
		return trueValueCallback()
	}
	return falseValue
}

// compare two ptr value is equal
func PtrValueEqual[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
