package server

import "log"

func (s *Server) initRoutes() {
	if s.router == nil {
		log.Fatal("routes init error: router isn't initialized")
	}

	s.router.HandleFunc("/api/request", s.handleGetRequests()).Methods("GET")
	s.router.HandleFunc("/api/requests", s.handleCreateRequest()).Methods("POST")

	log.Println("Server routes was initialized")
}
