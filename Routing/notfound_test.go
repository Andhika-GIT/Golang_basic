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

func TestNotFound(t *testing.T) {
	r := chi.NewRouter()

	r.NotFound(func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "gk ketemu")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	response := recorder.Result()

	resBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, "gk ketemu", string(resBody))
}
