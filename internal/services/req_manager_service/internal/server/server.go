package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	database *sql.DB
}

func New(dataBase *sql.DB) *Server {
	s := Server{
		router:   mux.NewRouter(),
		database: dataBase,
	}

	s.initRoutes()

	log.Println("Server instance created")
	return &s
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(":8000", s.router))
}
