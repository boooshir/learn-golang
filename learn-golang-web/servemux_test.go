package learngolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Dunia")
	})
	mux.HandleFunc("/hai", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hai dunia")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hai images")
	})
	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hai thumbnail")
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

func TestRequest(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
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
