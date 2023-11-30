package routing

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		panic("foo")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	response := recorder.Result()

	resBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "hello world", string(resBody))
}
