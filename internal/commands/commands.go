package commands

import (
	"fmt"
	"log/slog"
	"zeabot/internal/music"

	"github.com/bwmarrin/discordgo"
)

type data struct {
	*music.MusicManager
}

var d *data = &data{music.NewMusicManager()}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	joinName:  d.onJoinCommand,
	leaveName: d.onLeaveCommand,
	pingName:  d.onPingCommand,
	playName:  d.onPlayCommand,
}

var commands = []*discordgo.ApplicationCommand{
	joinCommand,
	leaveCommand,
	pingCommand,
	playCommand,
}

func RegisterCommands(s *discordgo.Session) func() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	var (
		registeredCommands []*discordgo.ApplicationCommand
		err                error
	)
	if registeredCommands, err = s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands); err != nil {
		slog.Error("Failed to register commands: ", slog.Any("err", err))
	}

	return func() {
		for _, command := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, "", command.ID)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to deregister '%v' command: %v", command.Name, err))
				continue
			}
		}
	}
}

type DifferentChannelError struct {
	s string
}

func NewDifferentChannelError(text string) error {
	return &DifferentChannelError{text}
}

func (e *DifferentChannelError) Error() string {
	return e.s
}

func assertVoiceConnection(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) (*discordgo.VoiceConnection, error) {
	var (
		vc              *discordgo.VoiceConnection
		g               *discordgo.Guild
		authorChannelId string
		ok              bool
		err             error
	)

	g, err = s.Guild(i.GuildID)
	if err != nil {
		slog.Error("Failed to guild: ", slog.Any("err", err))
	}

	for _, vs := range g.VoiceStates {
		if vs.UserID == i.Member.User.ID {
			authorChannelId = vs.ChannelID
		}
	}

	if vc, ok = s.VoiceConnections[i.GuildID]; !ok {
		vc, err = s.ChannelVoiceJoin(i.GuildID, authorChannelId, false, true)
		if err != nil {
			return nil, err
		}
	}

	if vc.ChannelID != authorChannelId {
		return nil, NewDifferentChannelError("Different channel")
	}

	return vc, nil
}
