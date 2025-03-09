package commands

import (
	"context"
	"fmt"

	"github.com/Zead0n/zeabot-go/zeabot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var leave = discord.SlashCommandCreate{
	Name:        "leave",
	Description: "Leave voice channel",
}

func (bot *botData) onLeave(
	command discord.SlashCommandInteractionData,
	event *handler.CommandEvent,
) error {
	bot.Lavalink.RemovePlayer(*event.GuildID())

	queue := bot.Manager.Get(*event.GuildID())
	queue.Clear()
	queue.Mode = zeabot.LoopOff

	if err := bot.Discord.UpdateVoiceState(context.TODO(), *event.GuildID(), nil, false, false); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("Error disconnecting: `%s`", err),
		})
	}

	return event.CreateMessage(discord.MessageCreate{
		Content: "Player Disconnected",
	})
}
