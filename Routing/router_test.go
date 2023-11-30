package routing

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var url string = "http://localhost:8080/"
var url2 string = "http://localhost:8080/product/1"

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
