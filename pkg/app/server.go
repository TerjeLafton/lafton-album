package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terjelafton/lafton-album/pkg/api"
)

type Server struct {
	router       *gin.Engine
	albumService api.AlbumService
}

func NewServer(router *gin.Engine, albumService api.AlbumService) *Server {
	return &Server{
		router:       router,
		albumService: albumService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("Server- there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
