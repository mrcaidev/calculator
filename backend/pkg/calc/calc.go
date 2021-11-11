// @Title		calc
// @Description	计算器组件。
// @Author		蔡与望
package calc

import (
	"backend/pkg/stack"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Pop out the top element as float64.
func asFloat(value interface{}) float64 {
	switch value := value.(type) {
	case float64:
		return value
	case string:
		topDouble, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err.Error())
		}
		return topDouble
	default:
		panic(fmt.Sprintf("popDouble() error: Invalid type %T", value))
	}
}

// Check if a string is a valid number or dot.
func isValidNumber(s string) bool {
	for _, char := range s {
		if (char < '0' || char > '9') && char != '.' {
			return false
		}
	}
	return true
}

// Check if a character is a valid operator. (self-defined operators included)
func isValidOperator(s string) bool {
	return strings.Contains("+-*/()sctpru!", s)
}

// Judge the priority of operator.
func opLevel(op string) int {
	switch op {
	case "(", ")":
		return 3
	case "s", "c", "t", "p", "r", "u", "!":
		return 2
	case "*", "/":
		return 1
	default:
		return 0
	}
}

// Compare the priority of two operators.
// Returns positive if a > b, negative if a < b, zero if a == b.
func opCmp(a string, b string) int {
	return opLevel(a) - opLevel(b)
}

// Calculate the factorial of the floor of a float number.
func factorial(x float64) float64 {
	if x == 0 {
		return 1.0
	}
	ret := 1.0
	for num := 1.0; num <= x; num++ {
		ret *= num
	}
	return ret
}

// Append a character to a slice, except parentheses.
func appendPart(dst []string, src string) []string {
	if src == "(" || src == ")" {
		return dst
	}
	return append(dst, src)
}

// Change the position of "p", "u" and "!".
func organize(exp string) string {
	var (
		targetIndex     int
		leftBracketNum  int
		rightBracketNum int
	)
	for index := 0; index < len(exp); index++ {
		targetIndex = 0
		leftBracketNum = 0
		rightBracketNum = 0
		newExpArr := make([]string, 0)
		if exp[index] == 'p' || exp[index] == 'u' || exp[index] == '!' {
			// Target the object to be operated on.
			if exp[index-1] == ')' {
				// Find corresponding left parenthesis.
				for rev := index - 1; rev >= 0; rev-- {
					if exp[rev] == ')' {
						rightBracketNum++
					} else if exp[rev] == '(' {
						leftBracketNum++
					}
					if leftBracketNum == rightBracketNum {
						targetIndex = rev
						break
					}
				}
			} else {
				// Find out where the operated number starts.
				for rev := index - 1; rev >= 0; rev-- {
					if !isValidNumber(string(exp[rev])) {
						targetIndex = rev + 1
						break
					}
				}
			}
			// Rearrange the expression.
			if targetIndex != 0 {
				newExpArr = append(newExpArr, exp[:targetIndex])
			}
			newExpArr = append(newExpArr, string(exp[index]))
			newExpArr = append(newExpArr, exp[targetIndex:index])
			if index != len(exp)-1 {
				newExpArr = append(newExpArr, exp[index+1:])
			}
			exp = strings.Join(newExpArr, "")
		}
	}
	return exp
}

// Fragment sorted string into number parts and operator parts.
func fragment(raw string) []string {
	parts := make([]string, 0)
	var part strings.Builder
	for _, charRune := range raw {
		char := string(charRune)
		if isValidNumber(char) {
			part.WriteString(char)
		} else if isValidOperator(char) {
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

// Convert original infix expression to postfix expression.
func toPostfix(infixArr []string) []string {
	postfixArr := make([]string, 0)
	opStack := stack.CreateStack()
	for _, curPart := range infixArr {
		if isValidNumber(curPart) {
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
		if isValidNumber(curPart) {
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
	rawExpression := organize(expression)
	parts := fragment(rawExpression)
	postfixArr := toPostfix(parts)
	rawAnswer := getAnswer(postfixArr)
	answer := strconv.FormatFloat(rawAnswer, 'f', -1, 32)
	return answer
}
