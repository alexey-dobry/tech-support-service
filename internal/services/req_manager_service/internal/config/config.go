package config

import (
	"log"

	_ "github.com/alexey-dobry/tech-support-platform/internal/pkg/config"
	"github.com/spf13/viper"
)

type Config struct {
	AuthServer AuthConfig `validate:"required" mapstructure:"server"`
}

type AuthConfig struct {
	MySqlDsn     string `validate:"required" mapstructure:"dsn"`
	ServerAdress string `validate:"required" mapstructure:"adress"`
}

func (cfg *Config) InitBotConfig() {
	cfg.AuthServer.MySqlDsn = viper.GetString("database.dsn")
	cfg.AuthServer.ServerAdress = viper.GetString("server.adress")
}

func Get() Config {
	cfg := Config{}

	cfg.InitBotConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Config Unmarshal error: %s", err)
	}

	return cfg
}
