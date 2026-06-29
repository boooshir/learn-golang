package learn

import "fmt"

func Slice() {
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jummat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "UPs"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "bangsyir"
	newSlice[1] = "ahmad"
	newSlice2 := append(newSlice, "test")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

}
