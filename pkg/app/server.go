package app

import (
	"log"
	"net/http"

	"github.com/terjelafton/lafton-album/pkg/api"
)

type Server struct {
	router       *Router
	albumService api.AlbumService
}

func NewServer(router *Router, albumService api.AlbumService) *Server {
	return &Server{
		router:       router,
		albumService: albumService,
	}
}

func (s *Server) Run(addr string) {
	r := s.Routes()

	log.Fatal(http.ListenAndServe(addr, r))
}
