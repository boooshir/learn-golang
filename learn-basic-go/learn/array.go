package learn

import "fmt"

func DataArray() {
	var names [3]string

	names[0] = "Eko"
	names[1] = "Kusniawan"
	names[2] = "Khannedy"
	fmt.Println(names)

	var values = [...]int{
		90, 80, 70, 100, 123,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[1:4])
}
