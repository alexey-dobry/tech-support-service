package app

import (
	"log"

	"database/sql"

	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/config"
	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/server"
)

type App struct {
	server *server.Server
}

func New(db *sql.DB) *App {
	a := App{
		server: server.New(db),
	}

	log.Print("App instance created")
	return &a
}

func (a *App) Run(cfg *config.Config) {
	log.Print("App is running...")
	a.server.Run(cfg)
}
