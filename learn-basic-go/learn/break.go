package learn

import "fmt"

func Break() {
	// for i := 0; i < 10; i++ {
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println("perulangan ke", i)
	// }
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("perulangan ke", i)
	}
}
