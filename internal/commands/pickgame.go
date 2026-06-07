package commands

import (
	"math/rand/v2"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var pickGameCommand = &discordgo.ApplicationCommand{
	Name:        "pick-game",
	Description: "escolhe um dos jogos aleatoriamente",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "jogos",
			Description: "Nomes dos jogos, separados por vírgula",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
	},
}

func parseGamesString(games string) []string {
	return strings.Split(games, ",")
}

func PickGameHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	if data.Name != "pick-game" {
		return
	}

	games := parseGamesString(data.Options[0].StringValue())

	pickedGame := games[rand.IntN(len(games))]

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "O jogo escolhido foi " + pickedGame,
		},
	})
}
