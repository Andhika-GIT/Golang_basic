package Server

import (
	"fmt"
	"net/http"
	"testing"
)

func setCookie(writer http.ResponseWriter, request *http.Request) {
	// create new cookie
	cookie := new(http.Cookie)

	// create new query
	parameter := request.URL.Query()

	cookie.Name = "Sponsored-by"          // cookie-key
	cookie.Value = parameter.Get("param") // cookie-value (set to parameter that we get from url)
	cookie.Path = "/"                     // url that's allow to access this cookie ("/" -> entire url page)

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "sucessfully set cookie")
}

func getCookie(writer http.ResponseWriter, request *http.Request) {
	// get the cookie from client
	cookie, err := request.Cookie("Sponsored-by")

	if err != nil {
		// if cookie empty
		fmt.Fprint(writer, "no cookie")
	} else {
		fmt.Fprintf(writer, "cookie value is %v", cookie.Value)
	}
}

// run server
func TestSetCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-cookie", getCookie)
	mux.HandleFunc("/set-cookie", setCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	// if server error
	if err != nil {
		panic(err)
	}
}
