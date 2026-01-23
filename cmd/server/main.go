package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TonChan8028/go-realtime-chat/internal/handler"
)

func main() {
	// ===== アプリケーション全体の Context =====
	ctx, cancel := context.WithCancel(context.Background())
	handler.SetAppContext(ctx)
	defer cancel()

	// ===== HTTP Server =====
	mux := http.NewServeMux()

	// WebSocket
	mux.HandleFunc("/ws", handler.WebSocketHandler)

	// 静的ファイル配信
	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fs)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// ===== OS Signal =====
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		log.Println("shutdown signal received")

		// 全体キャンセル
		cancel()

		// タイムアウト付き Shutdown
		ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelTimeout()

		if err := server.Shutdown(ctxTimeout); err != nil {
			log.Println("server shutdown error:", err)
		}
	}()

	log.Println("server started at :8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	log.Println("server exited")
}
