package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Zead0n/zeabot-go/commands"
	"github.com/Zead0n/zeabot-go/zeabot"
)

var err error

func main() {
	zeabot := zeabot.NewZeabot()

	zeabot.Discord.Rest().SetGlobalCommands(zeabot.Discord.ApplicationID(), commands.Commands)

	if err = zeabot.Discord.OpenGateway(context.TODO()); err != nil {
		slog.Error("Failed connecting gateway", slog.Any("err", err))
	}
}
