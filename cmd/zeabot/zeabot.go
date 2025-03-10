package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	defer func() {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	}()

	slog.Info("Bot started")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
	slog.Info("Bot shutting down")
}
