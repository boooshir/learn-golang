package learn

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil")
	message := recover()
	fmt.Println(message)
}

func runApplication(error bool) {
	defer logging()
	if error {
		panic("Ups error")
	}
	fmt.Println("Run application")
}

func Defer() {
	runApplication(true)
}
