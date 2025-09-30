package discord

import "github.com/bwmarrin/discordgo"

func (b *Bot) JoinVoiceChannel(guildId, channelId string) (*discordgo.VoiceConnection, error) {
	return b.Session.ChannelVoiceJoin(guildId, channelId, true, false)
}
