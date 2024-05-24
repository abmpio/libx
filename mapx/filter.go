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
