package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("firstname")
	lastName := request.PostForm.Get("lastname")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestHllo(t *testing.T) {
	requestBody := strings.NewReader("firstname=yusuf&lastname=ahmad")
	request := httptest.NewRequest("POST", "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
