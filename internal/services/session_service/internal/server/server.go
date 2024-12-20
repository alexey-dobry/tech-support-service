package server

import (
	"database/sql"
	"log"

	"github.com/alexey-dobry/tech-support-platform/internal/services/req_user_service/internal/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	database *sql.DB
}

func New(dataBase *sql.DB) *Server {
	s := Server{
		router:   gin.Default(),
		database: dataBase,
	}

	s.initRoutes()

	log.Println("Server instance created")
	return &s
}

func (s *Server) Run(cfg *config.Config) {
	log.Fatal(s.router.Run(":8070"))
}
