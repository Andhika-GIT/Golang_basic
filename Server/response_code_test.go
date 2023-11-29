package Server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func responseCodeHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(400) // badrequest
		fmt.Fprint(writer, "name parameter is empty")
	} else {
		writer.WriteHeader(200) // success
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	// url request
	// url := "http://localhost:8080/"
	url := "http://localhost:8080?name=andhika"

	request := httptest.NewRequest(http.MethodGet, url, nil)
	recorder := httptest.NewRecorder()

	responseCodeHandler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println("code :", response.StatusCode)
	fmt.Println(string(body))

}
