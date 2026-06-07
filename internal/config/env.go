package config

import "os"

type Config struct {
	DiscordToken  string
	ApplicationID string
	ServerID      string
}

func Load() (*Config, error) {
	cfg := &Config{
		DiscordToken:  os.Getenv("DISCORD_TOKEN"),
		ApplicationID: os.Getenv("APPLICATION_ID"),
		ServerID:      os.Getenv("SERVER_ID"),
	}

	return cfg, nil
}
