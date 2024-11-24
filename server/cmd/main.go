package main

import (
	"log"
	"tech-support-platform/internal/app"
	"tech-support-platform/internal/db"
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
