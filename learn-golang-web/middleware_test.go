package learngolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (ErrorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()
	ErrorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddlewareServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic executed")
		fmt.Fprintf(writer, "Hello Middleware")
		panic("ups")
	})
	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}
	errorhandler := &ErrorHandler{
		Handler: LogMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorhandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
