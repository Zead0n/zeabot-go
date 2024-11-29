package zeabot

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgolink/v3/disgolink"
)

var (
	discordToken     string = os.Getenv("DISCORD_TOKEN")
	lavalinkPassword string = os.Getenv("LAVALINK_PASSWORD")
)

type LoopState int64

const (
	LoopOff LoopState = iota
	LoopTrack
	LoopQueue
)

type Zeabot struct {
	Discord   bot.Client
	Lavalink  disgolink.Client
	LoopState LoopState
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

	lavalinkClient := disgolink.New(disgoClient.ApplicationID())
	node := disgolink.NodeConfig{
		Name:     "zeabot",
		Address:  "lavalink:2333",
		Password: lavalinkPassword,
		Secure:   true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	lavalinkClient.AddNode(ctx, node)

	zeabot.Discord = disgoClient
	zeabot.Lavalink = lavalinkClient
	zeabot.LoopState = LoopOff
	return zeabot
}
