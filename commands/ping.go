package commands

import "github.com/bwmarrin/discordgo"

var pingName string = "ping"

var pingCommand = &discordgo.ApplicationCommand{
	Name:        pingName,
	Description: "Just a pong",
}

func onPingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})
}
