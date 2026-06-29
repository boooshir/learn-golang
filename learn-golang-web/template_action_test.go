package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.html"))
	t.ExecuteTemplate(writer, "if.html", map[string]interface{}{
		// "Name": "",
	})
}

func TestTemplateIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
