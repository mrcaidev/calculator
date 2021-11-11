// @Title		calc
// @Description	计算器组件。
// @Author		蔡与望
package calc

import (
	"backend/pkg/stack"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// 将字符串或浮点数转换为浮点数。
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

// 判断字符串内容是否为浮点数。
func isNum(str string) bool {
	res, err := regexp.MatchString("^[\\d\\.]+$", str)
	if err != nil {
		panic(fmt.Sprintf("regexp.MatchString() error: %s.\n", err.Error()))
	}
	return res
}

// 判断字符内容是否为运算符。
func isOp(str string) bool {
	res, err := regexp.MatchString("^[\\+|\\-|\\*|/|\\(|\\)|s|c|t|p|r|u|!]$", str)
	if err != nil {
		panic(fmt.Sprintf("regexp.MatchString() error: %s.\n", err.Error()))
	}
	return res
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

// 拼接字符串，除了括号。
func appendPart(dst []string, src string) []string {
	if src == "(" || src == ")" {
		return dst
	}
	return append(dst, src)
}

// 找到与右括号平级的左括号。
func findLeftParen(parts []string, curIndex int) int {
	leftNum, rightNum := 0, 0
	for rev := curIndex; rev >= 0; rev-- {
		if parts[rev] == ")" {
			rightNum++
		} else if parts[rev] == "(" {
			leftNum++
		}
		if leftNum == rightNum {
			return rev
		}
	}
	return -1
}

// 将字符串分割为数字与运算符。
func fragment(str string) []string {
	parts := make([]string, 0)
	var part strings.Builder
	for _, charRune := range str {
		char := string(charRune)
		if isNum(char) {
			part.WriteString(char)
		} else if isOp(char) {
			if part.Len() > 0 {
				parts = append(parts, part.String())
				part.Reset()
			}
			parts = append(parts, char)
		} else {
			panic("Invalid character: " + char)
		}
	}
	if part.Len() > 0 {
		parts = append(parts, part.String())
	}
	return parts
}

// 调整特殊运算符位置。
func adjustOp(parts []string) []string {
	var targetIndex int
	for index := 0; index < len(parts); index++ {
		if parts[index] == "p" || parts[index] == "u" || parts[index] == "!" {
			if parts[index-1] == ")" {
				targetIndex = findLeftParen(parts, index-1)
			} else {
				targetIndex = index - 1
			}
			parts = append(parts[:targetIndex], append([]string{parts[index]}, parts[targetIndex:]...)...)[:len(parts)]
		}
	}
	return parts
}

// Convert original infix expression to postfix expression.
func toPostfix(infixArr []string) []string {
	postfixArr := make([]string, 0)
	opStack := stack.CreateStack()
	for _, curPart := range infixArr {
		if isNum(curPart) {
			// Directly append.
			postfixArr = appendPart(postfixArr, curPart)
		} else if curPart == "(" || opStack.Depth() == 0 {
			// Directly push.
			opStack.Push(curPart)
		} else if curPart == ")" {
			// Pop out and append every part between parentheses.
			for opStack.Depth() > 0 {
				top := opStack.Pop().(string)
				if top == "(" {
					break
				}
				postfixArr = appendPart(postfixArr, top)
			}
		} else {
			// Pop out and append higher-level operators before it,
			// till left parenthesis if there is one.
			for opStack.Depth() > 0 {
				topOp := opStack.Top().(string)
				if topOp == "(" || opCmp(curPart, topOp) > 0 {
					break
				}
				postfixArr = appendPart(postfixArr, opStack.Pop().(string))
			}
			// Push it into stack.
			opStack.Push(curPart)
		}
	}
	// Clear the stack.
	for opStack.Depth() > 0 {
		postfixArr = appendPart(postfixArr, opStack.Pop().(string))
	}
	return postfixArr
}

// Actual calculation using postfix string array.
func getAnswer(postfixArr []string) float64 {
	stack := stack.CreateStack()
	for _, curPart := range postfixArr {
		if isNum(curPart) {
			stack.Push(curPart)
			continue
		}
		res, top1, top2 := 0.0, 0.0, 0.0
		switch curPart {
		case "+":
			top1 = asFloat(stack.Pop())
			top2 = asFloat(stack.Pop())
			res = top1 + top2
		case "-":
			top1 = asFloat(stack.Pop())
			top2 = asFloat(stack.Pop())
			res = top2 - top1
		case "*":
			top1 = asFloat(stack.Pop())
			top2 = asFloat(stack.Pop())
			res = top1 * top2
		case "/":
			top1 = asFloat(stack.Pop())
			top2 = asFloat(stack.Pop())
			res = top2 / top1
		case "s":
			top1 = asFloat(stack.Pop())
			res = math.Sin(top1 * math.Pi / 180.0)
		case "c":
			top1 = asFloat(stack.Pop())
			res = math.Cos(top1 * math.Pi / 180.0)
		case "t":
			top1 = asFloat(stack.Pop())
			res = math.Tan(top1 * math.Pi / 180.0)
		case "p":
			top1 = asFloat(stack.Pop())
			res = math.Pow(top1, 2)
		case "r":
			top1 = asFloat(stack.Pop())
			res = math.Sqrt(top1)
		case "u":
			top1 = asFloat(stack.Pop())
			res = 1 / top1
		case "!":
			top1 = asFloat(stack.Pop())
			res = factorial(top1)
		}
		stack.Push(res)
	}
	return asFloat(stack.Pop())
}

// Calculate the answer of an expression.
func Calculate(expression string) string {
	parts := fragment(expression)
	parts = adjustOp(parts)
	postfixArr := toPostfix(parts)
	rawAnswer := getAnswer(postfixArr)
	answer := strconv.FormatFloat(rawAnswer, 'f', -1, 32)
	return answer
}
