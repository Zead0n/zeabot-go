package commands

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var joinName string = "join"

var joinCommand = &discordgo.ApplicationCommand{
	Name:        joinName,
	Description: "Just a pong",
}

func (d *data) onJoinCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, err := s.State.Guild(i.GuildID)
	if err != nil {
		if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "ERROR",
			},
		}); err != nil {
			slog.Error("Failed to send joined message: ", slog.Any("err", err))
		}
		slog.Error("Failed to guild: ", slog.Any("err", err))
		return
	}

	authorVoiceChannelId := ""
	for _, vs := range g.VoiceStates {
		if vs.UserID == i.Member.User.ID {
			authorVoiceChannelId = vs.ChannelID
			break
		}
	}

	if authorVoiceChannelId == "" {
		if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Couldn't find your voice channel",
			},
		}); err != nil {
			slog.Error("Failed to send joined message: ", slog.Any("err", err))
			return
		}
		return
	}

	if _, err = s.ChannelVoiceJoin(i.GuildID, authorVoiceChannelId, false, true); err != nil {
		slog.Error("Failed to join voice channel: ", slog.Any("err", err))
		return
	}

	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "joined",
		},
	}); err != nil {
		slog.Error("Failed to send joined message: ", slog.Any("err", err))
		return
	}
}
