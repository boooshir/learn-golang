package main

import "fmt"

func main() {
	x := 1
	y := 2
	ghostNumber := 29
	var pointer = &ghostNumber
	fmt.Printf("this is ghost number %v\n", &ghostNumber)
	fmt.Printf("this is pointer %v\n", pointer)

	fmt.Printf("x is %v and y is %v\n", x, y)

	fmt.Printf("this is pointer x is %v\n", &x)
	fmt.Printf("this is pointer y is %v\n", &y)
	// if run thi code below, it not gonna change the x and y
	// it just changing the copy code (memory location)
	// swap(x, y)
	// use reference instead to initiade x and y value (&)

	//passing the memmory pointer to the memory address
	// taking the reference
	swap(&x, &y) // example memoryloc 0x123 and 0x124
	fmt.Printf("x is %v and y is %v\n", x, y)
}

// *int = * here mean dereferences the memory loc 0x123 and 0x124
func swap(x, y *int) {
	var temp int = *x
	*x = *y
	*y = temp

	fmt.Printf("this is inside swap x is %v and y is %v\n", x, y)
}
