package learn

import "fmt"

func Map() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "arief"
	// person["address"] = "selayang"

	person := map[string]string{
		"name":    "arif",
		"address": "selayang",
	}
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "buku golang"
	book["author"] = "Sandi"
	book["price"] = "4_000_000"

	fmt.Println(book)
	delete(book, "price")
	fmt.Println(book)
}
