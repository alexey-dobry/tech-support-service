package server

import "log"

func (s *Server) initRoutes() {
	if s.router == nil {
		log.Fatal("Server router is not initialized")
	}

	s.router.POST("/auth", s.handleGetLoginData())

	log.Print("server routes was initialized")
}
