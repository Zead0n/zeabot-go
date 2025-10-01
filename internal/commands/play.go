package commands

import (
	"log/slog"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var playName string = "play"

var playCommand = &discordgo.ApplicationCommand{
	Name:        playName,
	Description: "Play some music",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "query",
			Description: "Query a track",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

var urlPattern = regexp.MustCompile(
	"^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?",
)

func (d *data) onPlayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, err := assertVoiceConnection(s, i)
	if err != nil {
		if err.Error() != "Different channel" {
			slog.Error("Error asserting voice connection: ", slog.Any("err", err))
			return
		}

		if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Already in a voice channel",
			},
		}); err != nil {
			slog.Error("Failed to send joined message: ", slog.Any("err", err))
			return
		}
	}
}
