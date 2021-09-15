// Package main instantiates all needed modules and starts the server.
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/terjelafton/lafton-album/pkg/api"
	"github.com/terjelafton/lafton-album/pkg/app"
	"github.com/terjelafton/lafton-album/pkg/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	if err := db.AutoMigrate(&api.Album{}); err != nil {
		log.Fatalf("Error running database migrations: %v", err)
	}

	storage := repository.NewStorage(db)
	albumService := api.NewAlbumService(storage)

	router := gin.Default()

	server := app.NewServer(router, albumService)

	if err := server.Run(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("AlbumRepository.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
