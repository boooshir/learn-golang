package learn

import "fmt"

func Ups() any {
  return 1
}

func Any() {
  var kosong any = Ups()
  fmt.Println("inteface kosong", kosong)

}
