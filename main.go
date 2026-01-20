package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/releases", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Release")
	})

	http.HandleFunc("/playlists", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Playlist")
	})

	http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Songs")
	})

	http.HandleFunc("/listeners", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Listener")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("error listening 8080")
	}
}
