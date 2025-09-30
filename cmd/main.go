package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"zeabot/internal/commands"

	"github.com/bwmarrin/discordgo"
)

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		slog.Error(err.Error())
		return
	}

	s.AddHandler(func(*discordgo.Session, *discordgo.Ready) {
		slog.Info(
			fmt.Sprintf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator),
		)
	})

	err = s.Open()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer s.Close()

	deregister := commands.RegisterCommands(s)
	defer deregister()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

// func onVoiceStateUpdate(s *discordgo.Session, event *discordgo.VoiceStateUpdate) {
// 	vc, ok := s.VoiceConnections[event.GuildID]
// 	if !ok {
// 		return
// 	}
//
// 	guild, err := s.Guild(vc.GuildID)
// 	if err != nil {
// 		slog.Error("Could not get guild on VoiceStateUpdate", slog.Any("err", err))
// 		return
// 	}
//
// 	var channelMembers []*discordgo.Member
// 	for _, state := range guild.VoiceStates {
// 		if state.ChannelID == vc.ChannelID && state.UserID != s.State.User.ID {
// 			channelMembers = append(channelMembers, state.Member)
// 		}
// 	}
//
// 	if len(channelMembers) > 1 {
// 		return
// 	}
//
// 	if err = vc.Disconnect(); err != nil {
// 		slog.Error("Failed to disconnect")
// 	}
// }
