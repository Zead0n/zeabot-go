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
					Value: zeabot.LoopOff.String(),
				},
				{
					Name:  "track",
					Value: zeabot.LoopTrack.String(),
				},
				{
					Name:  "queue",
					Value: zeabot.LoopQueue.String(),
				},
			},
		},
	},
}

func (data *botData) onLoop(
	command discord.SlashCommandInteractionData,
	event *handler.CommandEvent,
) error {
	player := data.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return event.CreateMessage(response.CreateWarn("No player exists"))
	}

	queueType := command.String("type")
	queue := data.Manager.Get(*event.GuildID())
	queue.SetLoop(queueType)

	return event.CreateMessage(response.Createf("Looping: %s", queueType))
}
