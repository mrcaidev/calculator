// @Title		calc
// @Description	计算器组件。
// @Author		蔡与望
package calc

import (
	"backend/packages/stack"
	"fmt"
	"math"
	"strings"
)

// 计算器内存。
type Calculator struct {
	parts []string
}

// 计算器初始化。
func (calc *Calculator) reset() {
	calc.parts = calc.parts[0:0]
}

// 找到与右括号平级的左括号。
func (calc *Calculator) findLeftParen(right int) int {
	leftNum, rightNum := 0, 0
	for rev := right; rev >= 0; rev-- {
		if calc.parts[rev] == ")" {
			rightNum++
		} else if calc.parts[rev] == "(" {
			leftNum++
		}
		if leftNum == rightNum {
			return rev
		}
	}
	return -1
}

// 将切片元素向前移动。
func (calc *Calculator) moveForward(src int, dst int) {
	value := calc.parts[src]
	copy(calc.parts[dst+1:src+1], calc.parts[dst:src])
	calc.parts[dst] = value
}

// 将表达式解析为数字与运算符。
func (calc *Calculator) parse(expression string) {
	var numPart strings.Builder
	// 逐个判断字符属于数字还是运算符。
	for _, charRune := range expression {
		char := string(charRune)
		if isNum(char) {
			numPart.WriteString(char)
		} else if isOp(char) {
			// 前面的数字结束了，添加进切片。
			if numPart.Len() > 0 {
				calc.parts = append(calc.parts, numPart.String())
				numPart.Reset()
			}
			// 当前的运算符也添加进切片。
			calc.parts = append(calc.parts, char)
		} else {
			panic(fmt.Sprintf("calc.fragment() error: Invalid character: %s.\n", char))
		}
	}
	// 最后剩下的数字也添加进切片。
	if numPart.Len() > 0 {
		calc.parts = append(calc.parts, numPart.String())
	}
}

// 将后置运算符前置。
func (calc *Calculator) adjustOp() {
	var target int
	for index := 0; index < len(calc.parts); index++ {
		if isPostOp(calc.parts[index]) {
			if calc.parts[index-1] == ")" {
				target = calc.findLeftParen(index - 1)
			} else {
				target = index - 1
			}
			calc.moveForward(index, target)
		}
	}
}

// 将中缀表达式转换为后缀表达式。
func (calc *Calculator) toPostfix() {
	postfixParts := make([]string, 0)
	opStack := stack.CreateStack()
	for _, curPart := range calc.parts {
		if isNum(curPart) {
			// 如果是数字，就直接加入表达式。
			postfixParts = append(postfixParts, curPart)
		} else if curPart == "(" || opStack.Depth() == 0 {
			// 如果运算符是“(”，或者栈空，就将运算符压入栈。
			opStack.Push(curPart)
		} else if curPart == ")" {
			// 如果运算符是“)”，就把两括号间的所有运算符弹出栈，并加入表达式。
			for opStack.Depth() > 0 {
				top := opStack.Pop().(string)
				if top == "(" {
					break
				}
				postfixParts = append(postfixParts, top)
			}
		} else {
			// 如果运算符不是括号，就弹出之前优先级更高的运算符，并加入表达式。
			// 直到碰到“(”或者低优先级运算符，这些不弹栈。
			for opStack.Depth() > 0 {
				top := opStack.Top().(string)
				if top == "(" || opCmp(curPart, top) > 0 {
					break
				}
				postfixParts = append(postfixParts, opStack.Pop().(string))
			}
			// 当前运算符压入栈。
			opStack.Push(curPart)
		}
	}
	// 清空栈。
	for opStack.Depth() > 0 {
		postfixParts = append(postfixParts, opStack.Pop().(string))
	}
	calc.parts = postfixParts
}

// 重构表达式各部分顺序。
func (calc *Calculator) refactor() {
	calc.adjustOp()
	calc.toPostfix()
}

// 计算答案。
func (calc *Calculator) getAnswer() string {
	stack := stack.CreateStack()
	for _, curPart := range calc.parts {
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
	return asString(asFloat(stack.Pop()))
}

// 运行计算器。
func (calc *Calculator) Calculate(expression string) string {
	calc.reset()
	calc.parse(expression)
	calc.refactor()
	return calc.getAnswer()
}
