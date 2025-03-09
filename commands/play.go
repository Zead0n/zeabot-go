package commands

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/Zead0n/zeabot-go/response"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
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

func (b *botData) onPlay(
	command discord.SlashCommandInteractionData,
	event *handler.CommandEvent,
) error {
	userVoiceState, ok := b.Discord.Caches().VoiceState(*event.GuildID(), event.User().ID)
	if !ok {
		return event.CreateMessage(discord.MessageCreate{
			Content: "Be in a voice channel",
			Flags:   discord.MessageFlagEphemeral,
		})
	}

	botVoiceState, ok := b.Discord.Caches().VoiceState(*event.GuildID(), event.ApplicationID())
	if !ok {
		if err := b.Discord.UpdateVoiceState(context.TODO(), *event.GuildID(), userVoiceState.ChannelID, false, true); err != nil {
			return event.CreateMessage(response.CreateErr("Error joining channel", err))
		}
	} else if userVoiceState.ChannelID == botVoiceState.ChannelID {
		// Check if the b is already in another channel
		return event.CreateMessage(response.CreateWarn("Already in another channnel"))
	}

	if err := event.DeferCreateMessage(false); err != nil {
		return err
	}

	query := command.String("query")
	if !urlPattern.MatchString(query) && !searchPattern.MatchString(query) {
		if source, ok := command.OptString("source"); ok {
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
			loadErr = b.handleTracks(event, track)
		},
		func(playlist lavalink.Playlist) {
			loadErr = b.handleTracks(event, playlist.Tracks...)
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

	player := b.Lavalink.Player(*event.GuildID())
	player.Node().LoadTracksHandler(ctx, query, resultHandler)

	return loadErr
}

func (b *botData) searchComponent(event *handler.CommandEvent, tracks ...lavalink.Track) error {
	trackIndexComponent := discord.NewTextInput(
		"trackIndex",
		discord.TextInputStyleShort,
		"Enter song number",
	)

	searchMessage := discord.NewMessageCreateBuilder().AddActionRow(trackIndexComponent)

	event.Client().Rest().CreateMessage(event.Channel().ID(), searchMessage.Build())
	return nil
}

func (b *botData) handleTracks(event *handler.CommandEvent, tracks ...lavalink.Track) error {
	b.AddTracks(*event.GuildID(), tracks...)

	var queuedMessage string
	for _, track := range tracks[1:min(len(tracks), 10)] {
		queuedMessage += fmt.Sprintf("Added to queue: %s\n", response.FormatTrack(&track))
	}

	_, err := event.UpdateInteractionResponse(response.Update(queuedMessage))
	if err != nil {
		_, err = event.UpdateInteractionResponse(
			response.UpdateErr("Queued song, but message failed", err),
		)
		return err
	}

	return nil
}
