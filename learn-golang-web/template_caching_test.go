package learngolangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.html
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.html", "Helo HTML Template")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
