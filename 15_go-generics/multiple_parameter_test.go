package gogenerics

import (
	"fmt"
	"testing"
)

func MultiplePrameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultiplePrameter("eko", 100)
	MultiplePrameter(100, "eko")
}
