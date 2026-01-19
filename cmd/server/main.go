package main

import (
	"log"
	"net/http"

	"github.com/TonChan8028/go-realtime-chat/internal/handler"
)

func main() {
	http.HandleFunc("/ws", handler.WebSocketEcho)

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
