package mapx

import "time"

func GetValueAsInt64Ptr(m map[string]interface{}, key string) *int64 {
	v, _ := GetValueAs[*int64](m, key)
	return v
}

func GetValueAsBoolPtr(m map[string]interface{}, key string) *bool {
	v, _ := GetValueAs[*bool](m, key)
	return v
}

func GetValueAsTimePtr(m map[string]interface{}, key string) *time.Time {
	v, _ := GetValueAs[*time.Time](m, key)
	return v
}

func GetValueAsString(m map[string]interface{}, key string) string {
	v, _ := GetValueAs[string](m, key)
	return v
}

func SetValue(m map[string]interface{}, key string, v interface{}, removeIfIsNil bool) {
	if m == nil {
		return
	}
	if removeIfIsNil && v == nil {
		delete(m, key)
	} else {
		m[key] = v
	}
}

// get a key value from map[string]interface{},
// T: try convert value to T, if cannot convert,return zero T value,and false
func GetValueAs[T any](m map[string]interface{}, key string) (T, bool) {
	var value T
	if len(m) <= 0 {
		return value, false
	}
	v, ok := m[key]
	if !ok {
		return value, false
	}
	convertedV, ok := v.(T)
	if !ok {
		return value, false
	}
	return convertedV, true
}
