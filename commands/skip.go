package commands

import (
	"context"

	"github.com/Zead0n/zeabot-go/response"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var skip = discord.SlashCommandCreate{
	Name:        "skip",
	Description: "Skip current track",
}

func (data *botData) onSkip(
	_ discord.SlashCommandInteractionData,
	event *handler.CommandEvent,
) error {
	player := data.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(response.CreateWarn("No player exists"))
	}

	queue := data.Manager.Get(*event.GuildID())

	nextTrack, _ := queue.Next()

	if err := player.Update(context.TODO(), lavalink.WithTrack(*nextTrack)); err != nil {
		return event.CreateMessage(response.CreateErr("Error skipping track", err))
	}

	return event.CreateMessage(response.Create("Skipped current track"))
}
