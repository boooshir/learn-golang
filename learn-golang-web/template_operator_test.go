package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func OperatorValueTemplate(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/operator.html"))
	t.ExecuteTemplate(writer, "operator.html", map[string]interface{}{
		"FinalValue": 65,
	})
}

func TestOperatorValueTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	OperatorValueTemplate(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
