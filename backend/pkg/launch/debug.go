// @Title		debug
// @Author		蔡与望
// @Description	以终端模式启动后端。
package launch

import (
	"backend/pkg/calc"
	"fmt"
)

// 在终端启动。
func Debug() {
	fmt.Println("------------------------------")
	var (
		app        calc.Calculator
		expression string
	)
	for {
		fmt.Print(">>> ")
		fmt.Scanln(&expression)
		answer := app.Calculate(expression)
		fmt.Printf("Answer: %s\n", answer)
	}
}
