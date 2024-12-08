package server

import "log"

func (s *Server) initRoutes() {
	if s.router == nil {
		log.Fatal("Server router is not initialized")
	}

	s.router.POST("/create", s.handleAddNewManager())
	s.router.POST("/assign", s.handleAssingnManager())
	s.router.POST("/end", s.handleEndSession())
	s.router.GET("/sessions/manager/:manager_id", s.handleGetManagerData())
	s.router.GET("/sessions/manager/:client_id", s.handleGetClientData())

	log.Print("server routes was initialized")
}
