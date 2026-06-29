package learn

import "fmt"

func Math() {
	var a uint16 = 10
	var b uint16 = 34
	var c uint16 = a + b
	fmt.Println(c)

	// augmented asignment
	a += 20
	fmt.Println(a)

	// unary operator
	a++
	fmt.Println(a)
}
