package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Coutry     string
	Postalcode string
}

type Customer struct {
	FirstName string
	LastName  string
	City      string
	Age       int
	marriage  bool
	Hobbies   []string
	Address   []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "abu",
		LastName:  "Daud",
		City:      "Kuala lumpur",
		Age:       30,
		marriage:  false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
