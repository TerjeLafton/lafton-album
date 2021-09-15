// Package app takes care of everything related to the server, including router and api handlers.
package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terjelafton/lafton-album/pkg/api"
)

// Server is the primary type that binds the router and the services together.
type Server struct {
	router       *gin.Engine
	albumService api.AlbumService
}

// NewServer is the constructor for the Server type and takes a gin engine as a router and an interface for each other service.
func NewServer(router *gin.Engine, albumService api.AlbumService) *Server {
	return &Server{
		router:       router,
		albumService: albumService,
	}
}

// Run is the entrypoint for the application and is used by main to start the server.
func (s *Server) Run() error {
	r := s.Routes()

	// Run the server through the gin engine router.
	if err := r.Run(); err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
