package discord

import (
	"zeabot/internal/music"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Session *discordgo.Session
	Music   *music.MusicManager
}

func NewBot(token string) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Session: s,
		Music:   music.NewMusicManager(),
	}

	return bot, nil
}

func (b *Bot) Init() {
	b.Session.Open()
}

func (b *Bot) Deinit() {
	b.Session.Close()
}
