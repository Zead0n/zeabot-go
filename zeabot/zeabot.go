package zeabot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

var (
	discordToken     string = os.Getenv("DISCORD_TOKEN")
	lavalinkHostname string = os.Getenv("LAVALINK_HOSTNAME")
	lavalinkPort     string = os.Getenv("LAVALINK_PORT")
	lavalinkPassword string = os.Getenv("LAVALINK_PASSWORD")
)

type Zeabot struct {
	Discord  bot.Client
	Lavalink disgolink.Client
	Manager  *QueueManager
}

func NewZeabot() *Zeabot {
	zeabot := &Zeabot{}

	gateways := gateway.WithIntents(
		gateway.IntentMessageContent,
		gateway.IntentGuilds,
		gateway.IntentGuildMembers,
		gateway.IntentGuildMessages,
		gateway.IntentGuildPresences,
		gateway.IntentGuildVoiceStates,
	)

	caches := cache.WithCaches(
		cache.FlagGuilds,
		cache.FlagChannels,
		cache.FlagVoiceStates,
	)

	disgoClient, err := disgo.New(
		discordToken,
		bot.WithGatewayConfigOpts(gateways),
		bot.WithCacheConfigOpts(caches),
		bot.WithEventListenerFunc(zeabot.onDiscordEvent),
	)
	if err != nil {
		slog.Error("Error building bot", slog.Any("err", err))
	}

	lavalinkClient := disgolink.New(
		disgoClient.ApplicationID(),
		disgolink.WithListenerFunc(zeabot.onTrackEnd),
	)
	node := disgolink.NodeConfig{
		Name:     "zeabot",
		Address:  fmt.Sprintf("%s:%s", lavalinkHostname, lavalinkPort),
		Password: lavalinkPassword,
		Secure:   false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	lavalinkClient.AddNode(ctx, node)

	zeabot.Discord = disgoClient
	zeabot.Lavalink = lavalinkClient
	zeabot.Manager = &QueueManager{
		queues: make(map[snowflake.ID]*Queue),
	}

	return zeabot
}

func (z *Zeabot) AddTracks(guildId snowflake.ID, tracks ...lavalink.Track) error {
	if len(tracks) <= 0 {
		return errors.New("No tracks to queue")
	}

	var trackLimit []lavalink.Track
	trackLength := len(tracks)
	switch {
	case trackLength >= 10: // NOTE: Limits the max loading to 10
		trackLimit = tracks[:10]
	default:
		trackLimit = tracks[:]
	}

	player := z.Lavalink.Player(guildId)
	queue := z.Manager.Get(guildId)

	queue.Add(trackLimit...)
	if player.Track() == nil && len(queue.Tracks) > 0 {
		nextTrack, ok := queue.Next()
		if !ok {
			return errors.New("Queued tracks but can't play")
		}

		player.Update(context.TODO(), lavalink.WithTrack(*nextTrack))
	}

	return nil
}
