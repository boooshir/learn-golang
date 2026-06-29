package learngolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func redirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello redirect")
}

func redirectFrom(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/hello", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/redirect_to", redirectTo)
	mux.HandleFunc("/redirect_from", redirectFrom)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
