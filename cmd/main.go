package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"zeabot/internal/commands"
	"zeabot/internal/discord"
)

func main() {
	zeabot, err := discord.NewBot(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		slog.Error("Failed to create bot", slog.Any("err", err))
		return
	}
	defer zeabot.Deinit()

	deregister := commands.RegisterCommands(zeabot)
	defer deregister()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
