package commands

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var leaveName string = "leave"

var leaveCommand = &discordgo.ApplicationCommand{
	Name:        leaveName,
	Description: "Leave the voice channel",
}

func (d *data) onLeaveCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	vc, ok := s.VoiceConnections[i.GuildID]
	if !ok {
		if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Already not in a channel",
			},
		}); err != nil {
			slog.Error("Failed to send joined message: ", slog.Any("err", err))
		}
		return
	}

	if err := vc.Disconnect(); err != nil {
		slog.Error("Failed to Disconnect voice connection: ", slog.Any("err", err))
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Leaving voice channel",
		},
	}); err != nil {
		slog.Error("Failed to send joined message: ", slog.Any("err", err))
		return
	}
}
