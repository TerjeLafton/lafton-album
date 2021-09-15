// Package api defines the services and their interfaces.
package api

import (
	"errors"
	"strconv"
	"strings"
)

// AlbumService is the interface that defines the required method for an album services.
type AlbumService interface {
	New(album Album) error
	All() ([]Album, error)
}

// AlbumRepository is the interface that defines the required methods for an album repository.
type AlbumRepository interface {
	CreateAlbum(Album) error
	GetAlbums() ([]Album, error)
}

type albumService struct {
	storage AlbumRepository
}

// NewAlbumService is the constructor of the album service type that takes a album repository and returns a new album services.
func NewAlbumService(albumRepo AlbumRepository) AlbumService {
	return &albumService{
		storage: albumRepo,
	}
}

// New takes an album as input, filters through it and ensures that everything is correct before posting it to the storage module.
func (a *albumService) New(album Album) error {
	if album.Name == "" {
		return errors.New("album service - name required")
	}

	if album.Artist == "" {
		return errors.New("album service - artist required")
	}

	if album.Year == 0 {
		return errors.New("album service - year required")
	}

	if len(strconv.Itoa(album.Year)) != 4 {
		return errors.New("album service - year should be four digits")
	}

	album.Name = strings.Title(album.Name)
	album.Artist = strings.Title(album.Artist)

	if err := a.storage.CreateAlbum(album); err != nil {
		return err
	}

	return nil
}

// All calls the GetAlbums on the storage service and returns the result.
func (a *albumService) All() ([]Album, error) {
	albums, err := a.storage.GetAlbums()
	if err != nil {
		return nil, err
	}

	return albums, nil
}
