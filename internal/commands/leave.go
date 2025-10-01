package commands

import (
	"github.com/bwmarrin/discordgo"
)

var leaveName string = "leave"

var leaveCommand = &discordgo.ApplicationCommand{
	Name:        leaveName,
	Description: "Leave the voice channel",
}

func (d *data) onLeaveCommand(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	vc, ok := s.VoiceConnections[i.GuildID]
	if !ok {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Already not in a channel",
			},
		})
	}

	if err := vc.Disconnect(); err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Leaving voice channel",
		},
	})
}
