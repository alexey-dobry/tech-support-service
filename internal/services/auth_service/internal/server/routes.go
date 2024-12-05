package server

import "log"

func (s *Server) initRoutes() {
	if s.router == nil {
		log.Fatal("Server router is not initialized")
	}

	s.router.POST("/auth", s.handleGetLoginData())
	s.router.GET("/reg", s.handlePostLoginData())

	log.Print("server routes was initialized")
}
