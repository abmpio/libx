package mathx

import "math"

// 由于标准库math.Round在做四舍五入时，不能指定小数位数
// 这个函数可以在四舍五入时指定小数位数
func Round(v float64, precision int) float64 {
	if precision < 1 {
		return math.Round(v)
	}

	powV := math.Pow10(precision)
	result := math.Round(v*powV) / powV
	return TruncateFloat64(result, precision)
}
