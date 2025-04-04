package routing

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var url string = "http://localhost:8080/"
var url2 string = "http://localhost:8080/product/1"
var url3 string = "http://localhost:8080/product/1/name/vast"
var url4 string = "http://localhost:8080/files/hello.txt"

func TestRouter(t *testing.T) {
	r := chi.NewRouter()

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello world")
	})

	request := httptest.NewRequest("GET", url, nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)
	response := recorder.Result()

	resBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "hello world", string(resBody))
}

func TestRouterParam(t *testing.T) {
	r := chi.NewRouter()

	r.Get("/product/{id}", func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		fmt.Fprintf(writer, "product %s", id)
	})

	request := httptest.NewRequest("GET", url2, nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)
	response := recorder.Result()

	resBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "product 1", string(resBody))

}

func TestRouterMultipleParam(t *testing.T) {
	r := chi.NewRouter()

	r.Get("/product/{id}/name/{name}", func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")
		name := chi.URLParam(request, "name")

		fmt.Fprintf(writer, "product id %s, product name %s", id, name)
	})

	request := httptest.NewRequest("GET", url3, nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	response := recorder.Result()

	resBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "product id 1, product name vast", string(resBody))

}

//go:embed resources
var resources embed.FS

func TestRouterServeFile(t *testing.T) {
	r := chi.NewRouter()

	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))
	r.Handle("/files/*", http.StripPrefix("/files", fileServer))

	request := httptest.NewRequest("GET", url4, nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	response := recorder.Result()

	resBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hubla httpRouter", string(resBody))
}
