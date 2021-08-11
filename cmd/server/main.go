package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/terjelafton/lafton-album/pkg/api"
	"github.com/terjelafton/lafton-album/pkg/app"
	"github.com/terjelafton/lafton-album/pkg/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
	}
}

func run() error {
	db, err := setupDatabase()
	if err != nil {
		return err
	}

	storage := repository.NewStorage(db)
	if err := storage.RunMigrations(); err != nil {
		return err
	}

	router := gin.Default()
	router.Use(cors.Default())

	albumService := api.NewAlbumService(storage)

	server := app.NewServer(router, albumService)

	err = server.Run()
	if err != nil {
		return err
	}

	return nil
}

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("album.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
