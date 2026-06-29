package library

import (
	"fmt"
	"strings"
)

func Strings() {
  fmt.Println(strings.Split("huhu haha", " "))
  fmt.Println(strings.Contains("seafood", "foo"))
  fmt.Println(strings.ToLower("Eko Kurniawan"))
  fmt.Println(strings.ToUpper("eko kurniawan"))
  fmt.Println(strings.Trim("  tu bagus string  "," "))
  fmt.Println(strings.ReplaceAll("eko kurniawan eko junaidi", "eko", "budi"))
}
