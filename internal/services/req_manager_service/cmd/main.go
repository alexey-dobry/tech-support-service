package main

import (
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/req_user_service/internal/app"
	"github.com/alexey-dobry/tech-support-platform/internal/services/req_user_service/internal/db"
)

func main() {

	Database, err := db.NewMySQL()
	if err != nil {
		log.Fatal("Error creating database")
	}
	defer Database.Close()

	App := app.New(Database)

	App.Run()

}
