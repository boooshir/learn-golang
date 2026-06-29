package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "test",
		LastName:  "kualiti",
		City:      "selayang",
		Age:       39,
		marriage:  true,
		Hobbies:   []string{"hello", "ini", "test"},
	}

	bytes, _ := json.Marshal(customer)
	data := string(bytes)
	fmt.Println(data)
	fmt.Println(customer.Hobbies[1])
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"test","LastName":"kualiti","City":"selayang","Age":39,"Hobbies":["hello","ini","test"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)

}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "test",
		LastName:  "kualiti",
		City:      "selayang",
		Age:       39,
		marriage:  true,
		Hobbies:   []string{"hello", "ini", "test"},
		Address: []Address{{
			Street:     "selayang baru",
			Coutry:     "Selangor",
			Postalcode: "68190",
		}},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {

	jsonString := `{"FirstName":"test","LastName":"kualiti","City":"selayang","Age":39,"Hobbies":["hello","ini","test"],"Address":[{"Street":"selayang baru","Coutry":"Selangor","Postalcode":"68190"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer.Address[0].Street)
}
