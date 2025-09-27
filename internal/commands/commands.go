package commands

import (
	"fmt"
	"log/slog"
	"zeabot/internal/discord"

	"github.com/bwmarrin/discordgo"
)

type data struct {
	*discord.Bot
}

func RegisterCommands(b *discord.Bot) func() {
	d := &data{b}

	var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		pingName: d.onPingCommand,
		playName: d.onPlayCommand,
	}

	var commands = []*discordgo.ApplicationCommand{
		pingCommand,
		playCommand,
	}

	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	var (
		registeredCommands []*discordgo.ApplicationCommand
		err                error
	)
	if registeredCommands, err = b.Session.ApplicationCommandBulkOverwrite(b.Session.State.User.ID, "", commands); err != nil {
		slog.Error("Failed to register commands: ", slog.Any("err", err))
	}

	return func() {
		for _, command := range registeredCommands {
			err := b.Session.ApplicationCommandDelete(b.Session.State.User.ID, "", command.ID)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to deregister '%v' command: %v", command.Name, err))
				continue
			}
		}
	}
}
