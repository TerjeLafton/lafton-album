package api

type AlbumService interface{}

type AlbumRepository interface{}

type albumService struct {
	storage AlbumRepository
}

func NewAlbumService(albumRepo AlbumRepository) AlbumService {
	return &albumService{
		storage: albumRepo,
	}
}
