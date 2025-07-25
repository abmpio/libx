package mapx

import (
	"reflect"
	"strconv"
)

// get key value as
func GetKeyValueAs[T any](maps map[string]interface{}, key string) T {
	var tValue T
	if len(maps) <= 0 {
		return tValue
	}

	value, ok := maps[key]
	if !ok {
		return tValue
	}

	// value is T type instance, return value
	if v, ok := value.(T); ok {
		return v
	}

	valType := reflect.TypeOf(value)
	zeroType := reflect.TypeOf(tValue)

	// value is ptr，T is struct, then
	if valType != nil && valType.Kind() == reflect.Ptr && zeroType == valType.Elem() {
		v := reflect.ValueOf(value)
		if !v.IsNil() {
			return v.Elem().Interface().(T)
		}
	}

	// value is struct instance, T is ptr ,then
	if zeroType != nil && zeroType.Kind() == reflect.Ptr && valType == zeroType.Elem() {
		ptr := reflect.New(valType)
		ptr.Elem().Set(reflect.ValueOf(value))
		return ptr.Interface().(T)
	}

	return tryNumericConversion[T](value)
}

// tryNumericConversion 尝试将基础数值类型互转
func tryNumericConversion[T any](value interface{}) T {
	var zero T
	if value == nil {
		return zero
	}

	val := reflect.ValueOf(value)
	vType := reflect.TypeOf(zero)

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		num := val.Convert(reflect.TypeOf(float64(0))).Float()
		target := reflect.New(vType).Elem()
		switch vType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			target.SetInt(int64(num))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			target.SetUint(uint64(num))
		case reflect.Float32, reflect.Float64:
			target.SetFloat(num)
		default:
			return zero
		}
		return target.Interface().(T)

	case reflect.String:
		// 如果目标是数字类型，尝试字符串解析
		if vType.Kind() >= reflect.Int && vType.Kind() <= reflect.Int64 {
			if i, err := strconv.ParseInt(val.String(), 10, 64); err == nil {
				target := reflect.New(vType).Elem()
				target.SetInt(i)
				return target.Interface().(T)
			}
		} else if vType.Kind() == reflect.Float64 || vType.Kind() == reflect.Float32 {
			if f, err := strconv.ParseFloat(val.String(), 64); err == nil {
				target := reflect.New(vType).Elem()
				target.SetFloat(f)
				return target.Interface().(T)
			}
		} else if vType.Kind() == reflect.Bool {
			if b, err := strconv.ParseBool(val.String()); err == nil {
				target := reflect.New(vType).Elem()
				target.SetBool(b)
				return target.Interface().(T)
			}
		}
	}

	return zero
}
