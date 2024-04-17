package math

import (
	"fmt"
	"strconv"
)

// 截取一个32位数的指定位数,如 3.14159265,截取2位则返回3.14
// 如果截取过程中出错,则返回0
func TruncateFloat64(v float64, pricePrecision int) float64 {
	truncated := fmt.Sprintf("%."+strconv.Itoa(pricePrecision)+"f", v)
	truncateValue, err := strconv.ParseFloat(truncated, 64)
	if err != nil {
		return 0
	}
	return truncateValue
}
