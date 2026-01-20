package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/release", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Release")
	})

	http.HandleFunc("/playlist", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Playlist")
	})

	http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Songs")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("error listening 8080")
	}
}
