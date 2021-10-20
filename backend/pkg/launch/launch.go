// @Title: launch
// @Description: launch the calculator backend in debug or server mode.
// @Author: Yuwang Cai
package launch

import (
	"backend/pkg/calc"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
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

// Handle the calculation sent from frontend.
func calcHandler(sock *websocket.Conn) {
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

// Server mode. Launch at localhost:1234.
func Server() {
	http.Handle("/", http.FileServer(http.Dir("..")))
	http.Handle("/calculator", websocket.Handler(calcHandler))
	fmt.Println("Server started.")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("ListenAndServe() error: ", err.Error())
	}
}
