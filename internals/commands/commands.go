package commands

import "github.com/bwmarrin/discordgo"

func RegisterCommandHandlers() map[string]func(event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error {
	return map[string]func(event *discordgo.InteractionCreate, data discordgo.ApplicationCommandInteractionData) error{}
}
