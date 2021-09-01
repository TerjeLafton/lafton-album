package main

import (
	"fmt"
	"os"

	"github.com/terjelafton/lafton-album/pkg/api"
	"github.com/terjelafton/lafton-album/pkg/app"
	"github.com/terjelafton/lafton-album/pkg/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := setupDatabase()
	if err != nil {
		return err
	}

	storage := repository.NewStorage(db)
	albumService := api.NewAlbumService(storage)
	router := app.NewRouter()
	server := app.NewServer(router, albumService)

	server.Run(":8080")

	return nil
}

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("AlbumRepository.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
