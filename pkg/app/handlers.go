package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terjelafton/lafton-album/pkg/api"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "album collection API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) NewAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newAlbum api.Album

		if err := c.ShouldBindJSON(&newAlbum); err != nil {
			log.Printf("handler error: %V", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		if err := s.albumService.New(newAlbum); err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new album created",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		albums, err := s.albumService.GetAll()
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, albums)
	}
}
