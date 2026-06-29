package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "satu",
		Name:     "sepatu",
		Price:    30,
		ImageURL: "test.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
func TestJSONTagDecode(t *testing.T) {
	jsonstring := `{"id":"satu","name":"sepatu","price":30,"image_url":"test.png"}`
	jsonBytes := []byte(jsonstring)
	product := &Product{}
	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.ImageURL)
}
