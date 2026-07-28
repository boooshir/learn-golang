package gogenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}
func TestBag(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}

	PrintBag(numbers)

	names := Bag[string]{"bashir", "yusuf", "ahmad"}
	fmt.Println(names)

	PrintBag(names)
}
