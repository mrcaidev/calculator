// @Title		server
// @Author		蔡与望
// @Description	以服务器模式启动后端。
package launch

import (
	"backend/packages/calc"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// 处理`/calculator`页面发来的请求。
func calcHandler(sock *websocket.Conn) {
	var (
		calc       calc.Calculator
		err        error
		expression string
		answer     string
	)
	for {
		// 接收。
		err = websocket.Message.Receive(sock, &expression)
		if err != nil {
			fmt.Printf("websocket.Message.Receive() error: %s.\n", err.Error())
			break
		}
		fmt.Printf("Expression: %s\n", expression)
		// 计算。
		answer = calc.Calculate(expression)
		fmt.Printf("Answer: %s\n", answer)
		// 发送。
		err = websocket.Message.Send(sock, answer)
		if err != nil {
			fmt.Printf("websocket.Message.Send() error: %s.\n", err.Error())
			break
		}
	}
}

// 在1234端口启动服务器。
func Server() {
	http.Handle("/", http.FileServer(http.Dir("..")))
	http.Handle("/calculator", websocket.Handler(calcHandler))
	fmt.Println("Server started.")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("http.ListenAndServe() error: ", err.Error())
	}
}
