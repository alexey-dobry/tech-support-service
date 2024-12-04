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

func (cfg *Config) InitBotConfig() {

	cfg.Bot.Tocken = viper.GetString("bot.tocken")
	cfg.Bot.Language = viper.GetString("bot.language")
}

func Get() Config {
	cfg := Config{}

	cfg.InitBotConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Config Unmarshal error: %s", err)
	}

	return cfg
}
