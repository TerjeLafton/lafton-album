// Package repository defines the storage interface and the different storage implementations.
package repository

import (
	"errors"

	"github.com/terjelafton/lafton-album/pkg/api"
	"gorm.io/gorm"
)

// Storage is the interface that defines all the required methods by the storage type.
type Storage interface {
	CreateAlbum(album api.Album) error
	GetAlbums() ([]api.Album, error)
}

type storage struct {
	db *gorm.DB
}

// NewStorage is the constructor of the storage type and returns a new storage object with the provided db.
func NewStorage(db *gorm.DB) Storage {
	return &storage{
		db: db,
	}
}

// CreateAlbum takes an album as input and creates the album in the database.
func (s *storage) CreateAlbum(album api.Album) error {
	// Check if the album already exists.
	tempAlbum := album
	result := s.db.First(&tempAlbum)
	if result.RowsAffected != 0 {
		return errors.New("storage service - album already exists")
	}

	if result = s.db.Create(&album); result.Error != nil {
		return result.Error
	}

	return nil
}

// GetAlbums gathers all albums from the database and returns them.
func (s *storage) GetAlbums() ([]api.Album, error) {
	var albums []api.Album

	if result := s.db.Find(&albums); result.Error != nil {
		return nil, result.Error
	}

	return albums, nil
}
