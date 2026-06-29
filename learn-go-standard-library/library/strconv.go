package library

import (
	"fmt"
	"strconv"
)

func Strconf() {
	result, err := strconv.ParseBool("salah")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}
	resultInt, err := strconv.Atoi("90909090")

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

  binary := strconv.FormatInt(999,2)
  fmt.Println(binary)

  var string string = strconv.Itoa(999)
  fmt.Println(string)
}
