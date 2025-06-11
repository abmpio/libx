package mapx

// filter in map[string]v
func FilterInMapString[v any](m map[string]v, fn func(string, v) bool) map[string]v {
	result := make(map[string]v, 0)
	for eachKey, eachV := range m {
		if fn != nil {
			if !fn(eachKey, eachV) {
				continue
			}
		}
		result[eachKey] = eachV
	}
	return result
}

// 从一个map中截取出指定的key值,返回一个新的map
func FilterByKeys[v any](m map[string]v, keys []string) map[string]v {
	result := make(map[string]v)
	for _, key := range keys {
		if value, exists := m[key]; exists {
			result[key] = value
		}
	}
	return result
}

// 从一个map中截取出指定的key值,并使用valueFn函数来为指定的key设置一个新的值，返回一个新的map
func FilterByKeysWith[v any](m map[string]v, keys []string, valueFn func(key string) v) map[string]v {
	result := make(map[string]v)
	for _, key := range keys {
		value, exists := m[key]
		if !exists {
			continue
		}
		currentV := value
		if valueFn != nil {
			// 使用函数值
			currentV = valueFn(key)
		}
		result[key] = currentV
	}
	return result
}
