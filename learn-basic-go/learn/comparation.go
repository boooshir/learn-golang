package learn

import "fmt"

func Comparation() {
	name1 := "eko"
	name2 := "dono"
	var compare = name1 == name2
	fmt.Println(compare)

	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusAbsensi && lulusNilaiAkhir
	fmt.Println("lulus", lulus)
}
