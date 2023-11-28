package Server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web

		fmt.Fprint(writer, "hello world")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello world")
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "this is hi page")
	})
	mux.HandleFunc("/images", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "this is images page")
	})

	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "this is thumbnail page")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web

		fmt.Fprintln(writer, "hello world")
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello world")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	helloHandler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	fmt.Println("status code : ", response.StatusCode)
	fmt.Println("body : ", bodyString)
}
