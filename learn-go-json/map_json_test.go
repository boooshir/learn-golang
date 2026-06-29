package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMap(t *testing.T) {
	jsonRequest := `{"id":"satu","name":"sepatu","price":30,"image_url":"test.png"}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
}

func TestJsonMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P032",
		"name":      "Appe macbook pro",
		"price":     300000,
		"image_url": "test.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
