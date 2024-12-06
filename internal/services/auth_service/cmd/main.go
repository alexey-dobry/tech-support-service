package main

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/app"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/config"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/db"
)

func main() {

	cfg := config.Get()

	Database, err := db.NewMySQL(&cfg)
	if err != nil {
		log.Fatal("Error creating database")
	}
	defer Database.Close()

	App := app.New(Database)

	App.Run(&cfg)
}
