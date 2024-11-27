package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Commands = []discord.ApplicationCommandCreate{
	ping,
}

func CommandListener() bot.ConfigOpt {
	h := handler.New()

	h.Command("/ping", onPing)

	return bot.WithEventListeners(h)
}
