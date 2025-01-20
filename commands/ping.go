package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var ping = discord.SlashCommandCreate{
	Name:        "ping",
	Description: "pong",
}

func (d *botData) onPing(
	command discord.SlashCommandInteractionData,
	e *handler.CommandEvent,
) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("pong").Build())
}
