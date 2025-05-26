package commands

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	pingName: onPingCommand,
}

var commands = []*discordgo.ApplicationCommand{
	pingCommand,
}

func RegisterCommands(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	for _, command := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
		if err != nil {
			slog.Error(fmt.Sprintf("Couldn't create '%v' command: %v", command.Name, err))
			continue
		}
	}
}
