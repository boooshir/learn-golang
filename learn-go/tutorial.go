package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var name string

	fmt.Println("Whats is your name ?")
	fmt.Scan(&name)
	fmt.Println(name)

}
