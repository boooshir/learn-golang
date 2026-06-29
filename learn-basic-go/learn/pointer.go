package learn

import "fmt"

type Address struct {
	City, State, Country string
}

func Pointer() {
	// details pass by value
	// address1 := Address{"Selayang", "Selangor", "Malaysia"}
	// address2 := address1

	// address2.City = "Gombak"

	// fmt.Println(address1) // value not changing
	//  fmt.Println(address2)

	// details pass by value
	var address1 Address = Address{"Selayang", "Selangor", "Malaysia"}
	var address2 *Address = &address1

	address2.City = "Gombak"

	fmt.Println(address1) // value is changing
	fmt.Println(address2)
}

func AsterisOperator() {
	// details pass by value
	address1 := Address{"Selayang", "Selangor", "Malaysia"}
	address2 := &address1

	address2.City = "Gombak"

	*address2 = Address{"Kuala lumpur", "WPKL", "Malaysia"}

	fmt.Println(address1) // value not changing
	fmt.Println(address2)

}

func NewFunction() {
  var alamat1 *Address = new(Address)
  var alamat2 *Address= alamat1
  alamat2.Country = "uganda"

  fmt.Println(alamat1)
  fmt.Println(alamat2)
}

// pointer function

func ChangeCountryName(address *Address) {
  address.Country = "Zimbabwe"
}
func PointerFunction() {
  address := Address{}
  ChangeCountryName(&address)
  fmt.Println(address)
}
