package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Zead0n/zeabot-go/commands"
	"github.com/Zead0n/zeabot-go/zeabot"
)

var err error

func main() {
	zeabot := zeabot.NewZeabot()

	zeabot.Discord.Rest().SetGlobalCommands(zeabot.Discord.ApplicationID(), commands.Commands)
	zeabot.Discord.AddEventListeners(commands.CommandListener(zeabot))

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		zeabot.Discord.Close(ctx)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = zeabot.Discord.OpenGateway(ctx); err != nil {
		slog.Error("Failed connecting gateway", slog.Any("err", err))
		os.Exit(-1)
	}

	slog.Info("Bot started")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
	slog.Info("Bot shutting down")
}
