package repository

import (
	"github.com/terjelafton/lafton-album/pkg/api"
	"gorm.io/gorm"
)

type Storage interface {
	RunMigrations() error
	NewAlbum(request api.Album) error
	GetAlbums() ([]api.Album, error)
}

type storage struct {
	db *gorm.DB
}

func (s *storage) RunMigrations() error {
	if err := s.db.AutoMigrate(&api.Album{}); err != nil {
		return err
	}

	return nil
}

func (s *storage) NewAlbum(request api.Album) error {
	if err := s.db.Create(&request); err != nil {
		return err.Error
	}

	return nil
}

func (s *storage) GetAlbums() ([]api.Album, error) {
	var albums []api.Album
	if err := s.db.Find(&albums); err.Error != nil {
		return nil, err.Error
	}
	return albums, nil

}

func NewStorage(db *gorm.DB) Storage {
	return &storage{
		db: db,
	}
}
