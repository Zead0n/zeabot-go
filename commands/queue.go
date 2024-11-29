package commands

import (
	"fmt"

	"github.com/Zead0n/zeabot-go/response"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/lavaqueue-plugin"
)

var queue = discord.SlashCommandCreate{
	Name:        "queue",
	Description: "Show current queue",
}

func (data *botData) onQueue(
	command discord.SlashCommandInteractionData,
	event *handler.CommandEvent,
) error {
	player := data.Lavalink.Player(*event.GuildID())
	queue, err := lavaqueue.GetQueue(event.Ctx, player.Node(), *event.GuildID())
	if err != nil {
		event.CreateMessage(response.CreateErr("Failed to get queue", err))
	}

	if len(queue.Tracks) <= 0 {
		return event.CreateMessage(response.Create("Nothing in the queue"))
	}

	content := fmt.Sprintf("Queue(%d):\n", len(queue.Tracks))
	currentTrack := player.Track()
	content += fmt.Sprintf(
		"Now playing: %s\n%s / %s\n\n",
		response.FormatTrack(currentTrack),
		player.Position().String(),
		currentTrack.Info.Length.String(),
	)

	for i, track := range queue.Tracks {
		line := fmt.Sprintf(
			"%d. [%s - %s](<%s>)\n",
			i+1,
			track.Info.Author,
			track.Info.Title,
			*track.Info.URI,
		)

		content += line
	}

	return event.CreateMessage(response.Create(content))
}
