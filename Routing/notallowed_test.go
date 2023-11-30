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

func TestNotAllowed(t *testing.T) {
	r := chi.NewRouter()

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "gk boleh")
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "POST REQUEST")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	response := recorder.Result()

	resBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "gk boleh", string(resBody))
}
