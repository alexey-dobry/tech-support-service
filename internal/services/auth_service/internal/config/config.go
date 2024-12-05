package config

import (
	"log"

	_ "github.com/alexey-dobry/tech-support-platform/internal/pkg/config"
	"github.com/spf13/viper"
)

type Config struct {
	AuthServeer AuthConfig `validate:"required" mapstructure:"bot"`
}

type AuthConfig struct {
	MySqlDsn     string `validate:"required" mapstructure:"dsn"`
	ServerAdress string `validate:"required" mapstructure:"adress"`
}

func (cfg *Config) InitBotConfig() {
	cfg.AuthServeer.MySqlDsn = viper.GetString("database.dsn")
	cfg.AuthServeer.ServerAdress = viper.GetString("server.adress")
}

func Get() Config {
	cfg := Config{}

	cfg.InitBotConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Config Unmarshal error: %s", err)
	}

	return cfg
}
