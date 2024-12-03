package app

import (
	"database/sql"
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/req_user_service/internal/server"
)

type App struct {
	server *server.Server
}

func New(db *sql.DB) *App {
	a := &App{
		server: server.New(db),
	}

	log.Println("App instance created")
	return a
}

func (a *App) Run() {
	log.Println("App running...")
	a.server.Run()
}
