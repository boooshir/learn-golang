package learngolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(writer, "name.html", map[string]interface{}{
		"Title": "Hello dunia",
		"Name":  "Eko Kurniawan",
		"Address": map[string]interface{}{
			"Street": "Belum tersedia ya",
		},
	})
}

func TestTemplateMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.html"))
	t.ExecuteTemplate(writer, "name.html", Page{
		Title: "Hello dunia",
		Name:  "Eko Kurniawan",
		Address: Address{
			Street: "Jakarta",
		},
	})
}

func TestTemplateStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
