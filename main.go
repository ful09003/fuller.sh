package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed content
	contentFS embed.FS
)

func main() {
	staticRoot, err := fs.Sub(contentFS, "content")
	if err != nil {
		log.Fatal("error")
	}
	http.Handle("/", http.FileServer(http.FS(staticRoot)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
