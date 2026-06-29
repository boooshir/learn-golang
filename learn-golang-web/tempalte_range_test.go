package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.html"))
	t.ExecuteTemplate(writer, "range.html", map[string]interface{}{
		"Hobbies": []string{
			// "Gaming", "Browsing", "Coding",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
