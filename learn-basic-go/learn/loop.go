package learn

import "fmt"

func Loop() {
	for counter := 1; counter <= 28; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	names := []string{"Arief", "adhmad", "khanedy"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(name[i])
	// }
	for _, name := range names {
		fmt.Println(name)
	}
}
