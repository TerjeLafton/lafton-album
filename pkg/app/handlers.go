package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/terjelafton/lafton-album/pkg/api"
)

func (s *Server) apiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := map[string]string{
			"status": "success",
			"data":   "album API is running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) handlePostAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newAlbum api.Album

		response := make(map[string]string)

		if err := c.ShouldBindJSON(&newAlbum); err != nil {
			log.Printf("handler error: %v", err)

			response["status"] = "failure"
			response["message"] = err.Error()

			c.JSON(http.StatusBadRequest, response)

			return
		}

		if err := s.albumService.New(newAlbum); err != nil {
			log.Printf("handler error: %v", err)

			response["status"] = "failure"
			response["message"] = err.Error()

			c.JSON(http.StatusBadRequest, response)

			return
		}

		response["status"] = "success"
		response["message"] = "album was successfully created"

		c.JSON(http.StatusCreated, response)
	}
}

func (s *Server) handleGetAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := make(map[string]string)

		albums, err := s.albumService.All()
		if err != nil {
			log.Printf("handler error: %v", err)

			response["status"] = "failure"
			response["message"] = err.Error()

			c.JSON(http.StatusBadRequest, response)

			return
		}

		c.JSON(http.StatusOK, albums)
	}
}
