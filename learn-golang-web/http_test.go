package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello dunia")
}

func TestHttpTest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recoreder := httptest.NewRecorder()

	HelloHandler(recoreder, request)

	response := recoreder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
