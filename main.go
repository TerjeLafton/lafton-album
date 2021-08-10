package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Album struct {
	Name   string
	Artist string
	Year   int
}

var albums = []Album{
	{Name: "Rumours", Artist: "Fleetwood Mac", Year: 1977},
	{Name: "Bad", Artist: "Michael Jackson", Year: 1987},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/albums", AlbumsHandler)
	r.HandleFunc("/albums/{album}", AlbumHandler)

	log.Fatal(http.ListenAndServe(":5000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my album collection!")
}

func AlbumsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		out, err := json.Marshal(albums)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprint(w, string(out))
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var album Album
		if err := json.Unmarshal(body, &album); err != nil {
			fmt.Println(err)
			return
		}

		albums = append(albums, album)
	}
}

func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, album := range albums {

		if album.Name == vars["album"] {
			out, err := json.Marshal(album)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Fprint(w, string(out))
			return
		}

	}
	fmt.Fprintf(w, "%s is not in the collection", vars["album"])
}
