package commands

import (
	"context"
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var leave = discord.SlashCommandCreate{
	Name:        "leave",
	Description: "Leave voice channel",
}

func (data *botData) onLeave(event *handler.CommandEvent) error {
	data.Lavalink.RemovePlayer(*event.GuildID())

	if err := data.Discord.UpdateVoiceState(context.TODO(), *event.GuildID(), nil, false, false); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("Error disconnecting: `%s`", err),
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Content: "Player Disconnected",
	})
}
