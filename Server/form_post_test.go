package Server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func responseBodyFormHandler(writer http.ResponseWriter, request *http.Request) {
	// parse body from request from client
	err := request.ParseForm()

	// check if request body is a form
	if err != nil {
		panic(err)
	}

	// parse the form body request to variabel
	firstname := request.PostForm.Get("firstname")
	lastname := request.PostForm.Get("lastname")

	fmt.Fprintf(writer, "firstname is %s, lastname is %s", firstname, lastname)
}

func TestBodyFormResponse(t *testing.T) {
	// url request
	url := "http://localhost:8080/"

	// json body
	requestBody := strings.NewReader("firstname=Andhika&lastname=Hubla")

	// make the request
	request := httptest.NewRequest(http.MethodPost, url, requestBody)

	// configurate request header to take the form data
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	// run the function
	responseBodyFormHandler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
