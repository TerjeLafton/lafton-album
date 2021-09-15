package app

import "github.com/gin-gonic/gin"

// Routes is the method that defines every route of the application and points to specific handlers.
func (s *Server) Routes() *gin.Engine {
	router := s.router

	// Group API endpoints together as v1 for futureproofing.
	v1 := router.Group("v1/api")
	{
		v1.GET("/status", s.apiStatus())
		v1.GET("/albums", s.handleGetAlbums())
		v1.POST("albums", s.handlePostAlbum())
	}

	return router
}
