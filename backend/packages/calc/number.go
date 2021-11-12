// @Title		number
// @Author		蔡与望
// @Description	计算器对浮点数的操作函数。
package calc

import (
	"fmt"
	"regexp"
	"strconv"
)

// 判断字符串内容是否为数字与小数点的组合。
func isNum(str string) bool {
	res, err := regexp.MatchString("^[\\d\\.]+$", str)
	if err != nil {
		panic(fmt.Sprintf("regexp.MatchString() error: %s.\n", err.Error()))
	}
	return res
}

// 将字符串转换为浮点数。
func asFloat(value interface{}) float64 {
	switch value := value.(type) {
	case float64:
		return value
	case string:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(fmt.Sprintf("strconv.ParseFloat() error: %s.\n", err.Error()))
		}
		return floatValue
	default:
		panic(fmt.Sprintf("calc.asFloat() error: Invalid type %T.", value))
	}
}

// 将浮点数转换为字符串。
func asString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 32)
}

// 计算浮点数的阶乘。
func factorial(num float64) float64 {
	if num < 2 {
		return 1.0
	}
	product := 1.0
	for iter := 2.0; iter <= num; iter++ {
		product *= iter
	}
	return product
}
