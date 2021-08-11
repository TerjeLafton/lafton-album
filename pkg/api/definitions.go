package api

import (
	"errors"
	"strings"
)

type Album struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Year   string `json:"year"`
}

func (a *albumService) New(album Album) error {
	if album.Name == "" {
		return errors.New("album service - name required")
	}

	if album.Artist == "" {
		return errors.New("album service - artist required")
	}

	if album.Year == "" {
		return errors.New("album service - year required")
	}

	album.Name = strings.ToLower(album.Name)
	album.Artist = strings.ToLower(album.Artist)

	if err := a.storage.NewAlbum(album); err != nil {
		return err
	}

	return nil
}

func (a *albumService) GetAll() ([]Album, error) {
	albums, err := a.storage.GetAlbums()
	if err != nil {
		return nil, err
	}

	return albums, nil
}
