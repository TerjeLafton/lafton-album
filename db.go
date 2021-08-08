package laftonalbum

type Album struct {
	Name   string
	Artist string
	Year   int
}

type AlbumStore struct {
	store []Album
}

func (a *AlbumStore) GetAlbums() []Album {
	var albums []Album
	albums = append(albums, a.store...)
	return albums
}

func (a *AlbumStore) GetAlbum(name string) Album {
	for _, album := range a.store {
		if album.Name == name {
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
