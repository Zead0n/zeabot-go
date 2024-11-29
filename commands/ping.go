package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var ping = discord.SlashCommandCreate{
	Name:        "ping",
	Description: "pong",
}

func (d *botData) onPing(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("pong").Build())
}
