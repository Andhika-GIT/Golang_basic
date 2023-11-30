package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		fmt.Println(err)
	}

}
