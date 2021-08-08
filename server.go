package laftonalbum

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	store AlbumStore
	http.Handler
}

const jsonContentType = "application/json"

func (s *Server) albumsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(s.store.GetAlbums())
}

func (s *Server) albumHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/albums/")
	s.fetchAlbumFromStore(w, name)
}

func (s *Server) fetchAlbumFromStore(w http.ResponseWriter, name string) {
	album := s.store.GetAlbum(name)
	fmt.Fprint(w, album)
}

func (s *Server) AddAlbumToStore(w http.ResponseWriter, album Album) {
	s.store.PostAlbum(album)
}

func NewServer(store AlbumStore) *Server {
	s := new(Server)

	s.store = store

	router := http.NewServeMux()
	router.Handle("/albums", http.HandlerFunc(s.albumsHandler))
	router.Handle("/albums/", http.HandlerFunc(s.albumHandler))

	s.Handler = router

	return s
}
