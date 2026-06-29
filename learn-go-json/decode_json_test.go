package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonRequest := `{"FirstName":"abu","LastName":"Daud","City":"Kuala lumpur","Age":30}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}
