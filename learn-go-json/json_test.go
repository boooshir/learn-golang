package learngojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	logJson("Eko")
	logJson(1)
	logJson(true)
	logJson([]string{"eko", "sudin", "sumiati"})
	// logJson({"hello", "ini", "ada"})
}
