package main

import (
	"log"
	"tech-support-platform/internal/db"
	"tech-support-platform/internal/server"
)

func main() {
	Database, _ := db.NewMySQL()

	App := server.New(Database)

	App.Run()
	log.Println("Server is running...")
}
