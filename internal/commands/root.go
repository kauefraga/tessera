package commands

import "github.com/bwmarrin/discordgo"

func GenerateCommandsList() []*discordgo.ApplicationCommand {
	var commands []*discordgo.ApplicationCommand

	commands = append(commands, pickGameCommand)

	return commands
}
