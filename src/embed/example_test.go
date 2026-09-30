package embed_test

import (
	"embed"
	"log"
	"net/http"
)

//go:embed internal/embedtest/testdata/*.txt
var content embed.FS

func Example() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
