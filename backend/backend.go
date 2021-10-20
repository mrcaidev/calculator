// @Title: backend
// @Description: launch the calculator backend in debug or server mode.
// @Author: Yuwang Cai, Fumiama
package backend

import (
	"fmt"

	"golang.org/x/net/websocket"

	"github.com/MrCaiDev/GoCalculator/backend/calc"
)

// Debug mode. I/O at terminal.
func Debug() {
	println("------------------------------")
	var expression string
	for {
		fmt.Print(">>> ")
		fmt.Scanln(&expression)
		answer := calc.Calculate(expression)
		fmt.Printf("Answer: %s\n", answer)
	}
}

// CalcHandler Handle the calculation sent from frontend.
func CalcHandler(sock *websocket.Conn) {
	var (
		err        error
		expression string
	)
	for {
		err = websocket.Message.Receive(sock, &expression)
		if err != nil {
			fmt.Printf("Receive() error: %s\n", err.Error())
			break
		}
		fmt.Printf("Expression: %s\n", expression)
		answer := calc.Calculate(expression)
		fmt.Printf("Answer: %s\n", answer)
		err = websocket.Message.Send(sock, answer)
		if err != nil {
			fmt.Printf("Send() error: %s\n", err.Error())
			break
		}
	}
}
