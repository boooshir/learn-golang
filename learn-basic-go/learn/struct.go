package learn

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func Stuct() {
	var eko Customer
	eko.Name = "eko kurniawan"
	eko.age = 12
	eko.Address = "Jakarta"

	fmt.Println(eko)

	joko := Customer{
		Name:    "joko",
		Address: "Gombak",
		age:     39,
	}

	fmt.Println(joko)

	joko.sayHello("samsul")
}
