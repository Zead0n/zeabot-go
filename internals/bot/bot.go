package bot

import (
	"log/slog"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/v3/disgolink"
)

var (
	discordToken     string = os.Getenv("DISCORD_TOKEN")
	lavalinkHost     string = os.Getenv("LAVALINK_HOST")
	lavalinkPassword string = os.Getenv("LAVALINK_PASSWORD")
)

type Bot struct {
	Session  *discordgo.Session
	Lavalink disgolink.Client
	Handler  struct {
		Commands   map[string]func(event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error
		Components map[string]func(event *discordgo.InteractionCreate, data discordgo.MessageComponentInteractionData) error
	}
}

func New() (*Bot, error) {
	var bot = &Bot{}

	discordSession, err := discordgo.New("")
	if err != nil {
		return nil, err
	}

	discordSession.Identify.Intents = discordgo.IntentGuildMembers | discordgo.IntentGuildMessages | discordgo.IntentGuildPresences | discordgo.IntentGuilds | discordgo.IntentGuildVoiceStates

	bot.Session = discordSession

	return bot, nil
}

func (b *Bot) onApplicationCommand(session *discordgo.Session, event *discordgo.InteractionCreate) {
	switch event.Type {
	case discordgo.InteractionApplicationCommand:
		data := event.ApplicationCommandData()
		commandHandler, ok := b.Handler.Commands[data.Name]
		if !ok {
			slog.Error("Unknown application command: ", data.Name)
			return
		}
		if err := commandHandler(event, data); err != nil {
			slog.Error("Error handling command: ", err.Error())
		}
	}
}
