package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func GetProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	text := "Product " + params.ByName("id")
	fmt.Fprint(writer, text)
}

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", GetProducts)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1", string(bytes))

}
