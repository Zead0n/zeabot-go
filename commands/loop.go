package commands

import (
	"github.com/Zead0n/zeabot-go/response"
	"github.com/Zead0n/zeabot-go/zeabot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var loop = discord.SlashCommandCreate{
	Name:        "loop",
	Description: "Change loop state",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "type",
			Description: "Loop type",
			Required:    true,
			Choices: []discord.ApplicationCommandOptionChoiceString{
				{
					Name:  "off",
					Value: string(zeabot.LoopOff),
				},
				{
					Name:  "track",
					Value: string(zeabot.LoopTrack),
				},
				{
					Name:  "queue",
					Value: string(zeabot.LoopQueue),
				},
			},
		},
	},
}

func (bot *botData) onLoop(
	command discord.SlashCommandInteractionData,
	event *handler.CommandEvent,
) error {
	player := bot.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(response.CreateWarn("No player exists"))
	}

	queueType := zeabot.LoopState(command.String("type"))
	queue := bot.Manager.Get(*event.GuildID())

	queue.Mode = queueType

	return event.CreateMessage(response.Createf("Looping: %s", queueType))
}
