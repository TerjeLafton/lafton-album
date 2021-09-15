package api

type Album struct {
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}
