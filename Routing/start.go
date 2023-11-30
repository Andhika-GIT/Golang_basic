package routing

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func startServer() {
	r := chi.NewRouter()

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		fmt.Println(err)
	}

}
