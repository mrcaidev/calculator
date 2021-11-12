// @Title		operator
// @Author		蔡与望
// @Description	计算器对运算符的操作函数。
package calc

import (
	"fmt"
	"regexp"
)

// 判断字符串内容是否为运算符。
// 运算符列表：[+ - * / ( ) s c t p r u !]
func isOp(str string) bool {
	res, err := regexp.MatchString("^[\\+|\\-|\\*|/|\\(|\\)|s|c|t|p|r|u|!]$", str)
	if err != nil {
		panic(fmt.Sprintf("regexp.MatchString() error: %s.\n", err.Error()))
	}
	return res
}

// 判断运算符是否为后置运算符。
// 后置运算符列表：[p u !]
func isPostOp(op string) bool {
	return op == "p" || op == "u" || op == "!"
}

// 判断运算符优先级。
func opLevel(op string) int {
	switch op {
	case "(", ")":
		return 100
	case "s", "c", "t", "p", "r", "u", "!":
		return 10
	case "*", "/":
		return 1
	default:
		return 0
	}
}

// 比较两运算符优先级。
// a > b -> 正数
// a < b -> 负数
// a = b -> 0
func opCmp(op1 string, op2 string) int {
	return opLevel(op1) - opLevel(op2)
}
