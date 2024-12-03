package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./configs/config.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")
	viper.SafeWriteConfig()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}
