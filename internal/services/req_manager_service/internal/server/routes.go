package server

import "log"

func (s *Server) initRoutes() {
	if s.router == nil {
		log.Fatal("Server router is not initialized")
	}

	s.router.POST("/session", s.handleGetSession())
	s.router.POST("/create_session", s.handleCreateSession())

	log.Print("server routes was initialized")
}
