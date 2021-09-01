package app

func (s *Server) Routes() *Router {
	router := s.router

	router.HandleFunc("GET", "/status", s.handleStatus())

	return router
}
