package config

import (
	"log"

	_ "github.com/alexey-dobry/tech-support-platform/internal/pkg/config"
	"github.com/spf13/viper"
)

type Config struct {
	Bot BotConfig `validate:"required" mapstructure:"bot"`
}

type BotConfig struct {
	Tocken   string `validate:"required" mapstructure:"token"`
	Language string `validate:"required" mapstructure:"language"`
}

func init() {
	viper.SetEnvPrefix("MBOX")

	viper.BindEnv("bot.token")
	log.Print(Get().Bot.Tocken)
	viper.SetDefault("bot.language", "ru-RU")

	viper.SetDefault("managerclient.target", "localhost:50051")

	viper.SetDefault("pmclient.target", "localhost:50052")
}

func Get() Config {
	cfg := Config{}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Config Unmarshal error: %s", err)
	}

	return cfg
}
