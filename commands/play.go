package commands

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Zead0n/zeabot-go/response"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/json"
	"github.com/disgoorg/lavaqueue-plugin"
)

var (
	urlPattern = regexp.MustCompile(
		"^https?://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]?",
	)
	searchPattern = regexp.MustCompile(`^(.{2})(search|isrc):(.+)`)
)

var play = discord.SlashCommandCreate{
	Name:        "play",
	Description: "Play a track via url or search",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "query",
			Description: "Enter a query.",
			Required:    true,
		},
		discord.ApplicationCommandOptionString{
			Name:        "source",
			Description: "Changes source for a search",
			Required:    false,
			Choices: []discord.ApplicationCommandOptionChoiceString{
				{
					Name:  "Youtube",
					Value: string(lavalink.SearchTypeYouTube),
				},
			},
		},
	},
}

func (data *botData) onPlay(event *handler.CommandEvent) error {
	userVoiceState, ok := data.Discord.Caches().VoiceState(*event.GuildID(), event.User().ID)
	if !ok {
		return event.CreateMessage(discord.MessageCreate{
			Content: "Be in a voice channel",
		})
	}

	botVoiceState, ok := data.Discord.Caches().VoiceState(*event.GuildID(), event.ApplicationID())
	if !ok {
		if err := data.Discord.UpdateVoiceState(context.TODO(), *event.GuildID(), userVoiceState.ChannelID, false, false); err != nil {
			return event.CreateMessage(discord.MessageCreate{
				Content: "Error joininng voice channel",
			})
		}
	}

	// Check if the bot is already in another channel
	if botVoiceState.ChannelID != nil && userVoiceState.ChannelID != botVoiceState.ChannelID {
		return event.CreateMessage(discord.MessageCreate{
			Content: "Already in another channel",
		})
	}

	if err := event.DeferCreateMessage(false); err != nil {
		return err
	}

	query := event.SlashCommandInteractionData().String("query")
	if !urlPattern.MatchString(query) && !searchPattern.MatchString(query) {
		if source, ok := event.SlashCommandInteractionData().OptString("source"); ok {
			query = lavalink.SearchType(source).Apply(query)
		} else {
			query = lavalink.SearchTypeYouTube.Apply(query)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var loadErr error
	resultHandler := disgolink.NewResultHandler(
		func(track lavalink.Track) {
			loadErr = data.handleTracks(event, track)
		},
		func(playlist lavalink.Playlist) {
			loadErr = data.handleTracks(event, playlist.Tracks...)
		},
		func(tracks []lavalink.Track) {
			_, loadErr = event.UpdateInteractionResponse(
				response.UpdateWarn("Search has yet to be Implemented"),
			)
		},
		func() {
			_, loadErr = event.UpdateInteractionResponse(
				response.UpdateWarnf("No result for `%s`", query),
			)
		},
		func(err error) {
			_, loadErr = event.UpdateInteractionResponse(
				response.UpdateErr("Error while querying", err),
			)
		},
	)

	player := data.Lavalink.Player(*event.GuildID())
	player.Node().LoadTracksHandler(ctx, query, resultHandler)

	return loadErr
}

func (data *botData) handleTracks(event *handler.CommandEvent, tracks ...lavalink.Track) error {
	if len(tracks) <= 0 {
		event.UpdateInteractionResponse(response.UpdateWarn("No tracks to queue"))
		return nil
	}

	queueTracks := make([]lavaqueue.QueueTrack, len(tracks))
	var queuedMessage []string

	for i, track := range tracks {
		queuedMessage = append(
			queuedMessage,
			fmt.Sprintf(
				"Added to queue: [%s - %s](<%s>)",
				track.Info.Author,
				track.Info.Title,
				*track.Info.URI,
			),
		)

		queueTracks[i] = lavaqueue.QueueTrack{
			Encoded:  track.Encoded,
			UserData: nil,
		}
	}

	player := data.Lavalink.Player(*event.GuildID())
	_, err := lavaqueue.AddQueueTracks(event.Ctx, player.Node(), *event.GuildID(), queueTracks)
	if err != nil {
		_, err = event.UpdateInteractionResponse(response.UpdateErr("Error queuing", err))
		return err
	}

	_, err = event.UpdateInteractionResponse(discord.MessageUpdate{
		Content: json.Ptr(strings.Join(queuedMessage, "\n")),
	})

	return err
}
