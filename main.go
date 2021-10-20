package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/MrCaiDev/GoCalculator/backend"
	"golang.org/x/net/websocket"
)

//go:embed frontend
var frontend embed.FS

func main() {
	if len(os.Args) == 2 {
		fsys, err := fs.Sub(frontend, "frontend")
		if err != nil {
			panic(err)
		}
		http.Handle("/", http.FileServer(http.FS(fsys)))
		http.Handle("/calculator", websocket.Handler(backend.CalcHandler))
		log.Fatal(http.ListenAndServe(os.Args[1], nil))
	} else {
		fmt.Println("Usage: ip:port")
	}
}
