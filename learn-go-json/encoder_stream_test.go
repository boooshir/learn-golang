package learngojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoderStream(t *testing.T) {
	writer, _ := os.Create("test.json")

	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "junaidi",
		LastName:  "Kasino",
		Age:       56,
	}
	err := encoder.Encode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
