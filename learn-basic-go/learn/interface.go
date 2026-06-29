package learn

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
  Name string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
  return animal.Name
}

func Interface() {
	person := Person{Name: "saidin"}
  animal := Animal{Name: "marsupilami"}
	SayHello(person)
  SayHello(animal)
}
