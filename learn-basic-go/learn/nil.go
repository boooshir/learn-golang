package learn

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func Nil() {
	data := NewMap("hello")
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
	}
}
