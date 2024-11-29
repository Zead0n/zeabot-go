package commands

import (
	"github.com/Zead0n/zeabot-go/zeabot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Commands = []discord.ApplicationCommandCreate{
	ping,
	join,
	leave,
	play,
	queue,
}

type botData struct {
	*zeabot.Zeabot
}

func CommandListener(z *zeabot.Zeabot) handler.Router {
	cmds := &botData{z}

	handler := handler.New()
	handler.Command("/ping", cmds.onPing)
	handler.Command("/join", cmds.onJoin)
	handler.Command("/leave", cmds.onLeave)
	handler.Command("/play", cmds.onPlay)
	handler.SlashCommand("/queue", cmds.onQueue)

	return handler
}
