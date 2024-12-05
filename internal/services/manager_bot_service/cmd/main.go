package main

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/app"
	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/bot"
	"github.com/alexey-dobry/tech-support-platform/internal/services/manager_bot_service/internal/config"
)

func main() {
	log.Print("Building manager_bot_service...")

	cfg := config.Get()

	log.Print("Config read succesfully")

	bot := bot.New(&cfg)

	log.Print("Bot built succesfully")

	app := app.New(bot)

	log.Print("App build succesfully")

	log.Print("Build complete, service is running...")
	app.Run()
}
