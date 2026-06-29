package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.html"))
	t.ExecuteTemplate(writer, "address.html", map[string]interface{}{
		"Title": "Hello Dunia",
		"Name":  "Kita rubah dunia",
		"Address": map[string]interface{}{
			"Street": "Jalan bahagia",
			"City":   "Kuala lumpur",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
