package zeabot

import (
	"log/slog"
	"os"

	"github.com/Zead0n/zeabot-go/commands"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	// "github.com/disgoorg/disgo/handler"
	// "github.com/disgoorg/disgolink/v3/disgolink"
)

var (
	token string = os.Getenv("DISCORD_TOKEN")
)

type Zeabot struct {
	Discord bot.Client
}

func NewZeabot() *Zeabot {
	gateways := gateway.WithIntents(
		gateway.IntentMessageContent,
		gateway.IntentGuilds,
		gateway.IntentGuildMembers,
		gateway.IntentGuildMessages,
		gateway.IntentGuildPresences,
		gateway.IntentGuildVoiceStates,
	)

	client, err := disgo.New(
		token,
		commands.CommandListener(),
		bot.WithGatewayConfigOpts(gateways),
	)
	if err != nil {
		slog.Error("Error building bot", slog.Any("err", err))
	}

	return &Zeabot{
		Discord: client,
	}
}
