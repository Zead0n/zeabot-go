package commands

import (
	"context"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var join = discord.SlashCommandCreate{
	Name:        "join",
	Description: "Join the voice channel you're in.",
}

func (bot *botData) onJoin(
	command discord.SlashCommandInteractionData,
	event *handler.CommandEvent,
) error {
	voiceState, ok := bot.Discord.Caches().VoiceState(*event.GuildID(), event.User().ID)
	if !ok {
		return event.CreateMessage(discord.MessageCreate{
			Content: "Be in a voice channel",
			Flags:   discord.MessageFlagEphemeral,
		})
	}

	if err := bot.Discord.UpdateVoiceState(
		context.TODO(),
		*event.GuildID(),
		voiceState.ChannelID,
		false,
		true,
	); err != nil {
		return event.CreateMessage(discord.MessageCreate{
			Content: "Something went wrong while joining.",
		})
	}

	return event.CreateMessage(
		discord.NewMessageCreateBuilder().SetContent("Join successful").Build(),
	)
}
