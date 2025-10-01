package commands

import (
	"github.com/bwmarrin/discordgo"
)

var joinName string = "join"

var joinCommand = &discordgo.ApplicationCommand{
	Name:        joinName,
	Description: "Just a pong",
}

func (d *data) onJoinCommand(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	g, err := s.State.Guild(i.GuildID)
	if err != nil {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "ERROR",
			},
		})
	}

	authorVoiceChannelId := ""
	for _, vs := range g.VoiceStates {
		if vs.UserID == i.Member.User.ID {
			authorVoiceChannelId = vs.ChannelID
			break
		}
	}

	if authorVoiceChannelId == "" {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Couldn't find your voice channel",
			},
		})
	}

	if _, err = s.ChannelVoiceJoin(i.GuildID, authorVoiceChannelId, false, true); err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "joined",
		},
	})
}
