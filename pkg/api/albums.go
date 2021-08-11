package api

type AlbumService interface {
	New(album Album) error
	GetAll() ([]Album, error)
}

type AlbumRepository interface {
	NewAlbum(Album) error
	GetAlbums() ([]Album, error)
}

type albumService struct {
	storage AlbumRepository
}

func NewAlbumService(albumRepo AlbumRepository) AlbumService {
	return &albumService{
		storage: albumRepo,
	}
}
