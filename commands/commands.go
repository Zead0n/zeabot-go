package commands

import (
	"github.com/Zead0n/zeabot-go/zeabot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var Commands = []discord.ApplicationCommandCreate{
	join,
	leave,
	loop,
	ping,
	play,
	queue,
	skip,
}

type botData struct {
	*zeabot.Zeabot
}

func CommandListener(z *zeabot.Zeabot) handler.Router {
	cmds := &botData{z}

	handler := handler.New()
	handler.Command("/join", cmds.onJoin)
	handler.Command("/leave", cmds.onLeave)
	handler.SlashCommand("/loop", cmds.onLoop)
	handler.Command("/ping", cmds.onPing)
	handler.Command("/play", cmds.onPlay)
	handler.SlashCommand("/queue", cmds.onQueue)
	handler.SlashCommand("/skip", cmds.onSkip)

	return handler
}
