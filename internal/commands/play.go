package commands

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var playName string = "play"

var playCommand = &discordgo.ApplicationCommand{
	Name:        playName,
	Description: "Play some music",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:     "query",
			Type:     discordgo.ApplicationCommandOptionString,
			Required: true,
		},
	},
}

var urlPattern = regexp.MustCompile(
	"^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?",
)

func (d *data) onPlayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "play",
		},
	})
}
