package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is emty")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hi %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080?name=", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
}
