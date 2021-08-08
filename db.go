package laftonalbum

type Album struct {
	name   string
	artist string
	year   int
}

type AlbumStore struct {
	store []Album
}

func (a *AlbumStore) GetAlbums() []Album {
	var albums []Album
	for _, album := range a.store {
		albums = append(albums, album)
	}
	return albums
}

func (a *AlbumStore) GetAlbum(name string) Album {
	for _, album := range a.store {
		if album.name == name {
			return album
		}
	}
	return Album{}
}

func (a *AlbumStore) PostAlbum(album Album) {
	a.store = append(a.store, album)
}

func NewAlbumStore() *AlbumStore {
	return &AlbumStore{[]Album{}}
}
