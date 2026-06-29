package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestHTTPQuery(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=yusuf", nil)

	SayHello(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println("This is body", string(body))
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "hello %s %s", firstName, lastName)
}

func TestHTTPMultipleQuery(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?first_name=yusuf&last_name=ahmad", nil)

	MultipleParameter(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

func MultipleParameterValues(write http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(write, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=yusuf&name=ahmad", nil)

	MultipleParameterValues(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}
