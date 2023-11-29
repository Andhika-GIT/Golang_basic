package Server

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")

	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// by default, it will access /resources/static/index.html, so if we enter /static/index.html in url, then it will cause error
	// so we use stripprefix to remove "/resources/", so we can access /static/index.html in url
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	// if server is error
	if err != nil {
		panic(err)
	}
}

// embed -> automatically embed our resources to binary distribution file, so when we hosted the app, we don't need to upload the statis resources folder again

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	// by default, it will access /static/resources/index.html, so if we enter /static/index.html in url, then it will cause error
	// so we use fs.sub, to get inside the "resources" folder, so we don't have to type "static/resources/index.html", instead we can use "static/index.html"
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	// if server is error
	if err != nil {
		panic(err)
	}

}
