package Server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
func helloNameHandler(writer http.ResponseWriter, request *http.Request) {
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")
	if firstname == "" || lastname == "" {
		fmt.Fprint(writer, "hello")
	} else {
		fmt.Fprintf(writer, "hello %s %s", firstname, lastname)
	}
}

func multipleParameterHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestQueryParameterHttp(t *testing.T) {
	// url1 := "http://localhost:8080/hello?first_name=andhika"
	// url2 := "http://localhost:8080/hello?first_name=andhika&last_name=pintar"
	url3 := "http://localhost:8080/hello?name=andhika&name=pintar"
	request := httptest.NewRequest(http.MethodGet, url3, nil)
	recorder := httptest.NewRecorder()

	multipleParameterHandler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)

	fmt.Println("status code : ", response.StatusCode)
	fmt.Println("body : ", bodyString)
}
