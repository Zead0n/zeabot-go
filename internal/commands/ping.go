package commands

import (
	"github.com/bwmarrin/discordgo"
)

var pingName string = "ping"

var pingCommand = &discordgo.ApplicationCommand{
	Name:        pingName,
	Description: "Just a pong",
}

func (d *data) onPingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	embed := &discordgo.MessageEmbed{
		Color:       0x4a90e2,
		Title:       "**INFO**",
		Description: "ping",
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}
