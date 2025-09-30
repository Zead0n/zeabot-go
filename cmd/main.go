package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"zeabot/internal/commands"

	"github.com/bwmarrin/discordgo"
)

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	s.AddHandler(func(*discordgo.Session, *discordgo.Ready) {
		slog.Info(
			fmt.Sprintf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator),
		)
	})

	err = s.Open()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer s.Close()

	commands.RegisterCommands(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	commands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to retrieve commands: %v", err))
		return
	}

	for _, command := range commands {
		err = s.ApplicationCommandDelete(s.State.User.ID, "", command.ID)
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to deregister '%v' command: %v", command.Name, err))
			continue
		}
	}
}
