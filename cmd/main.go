package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/kauefraga/tessera/internal/commands"
	"github.com/kauefraga/tessera/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("Failed to load config")
	}

	discord, err := discordgo.New(cfg.DiscordToken)

	discord.AddHandler(commands.PickGameHandler)

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})

	cmds := commands.GenerateCommandsList()

	_, err = discord.ApplicationCommandBulkOverwrite(cfg.ApplicationID, cfg.ServerID, cmds)
	if err != nil {
		log.Fatalln("Failed to register commands")
	}

	err = discord.Open()
	if err != nil {
		log.Fatalln("Failed to open session")
	}

	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Fatalln("Gracefully shutting down.")
}
