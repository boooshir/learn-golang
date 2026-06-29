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

func TestPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemid := params.ByName("itemId")

		text := "Product " + id + " Item " + itemid
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 1 Item 1", string(bytes))

}

func TestCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")

		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image : /small/profile.png", string(bytes))

}
