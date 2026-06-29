package learn

import "fmt"

type Man struct {
  Name string 
}

func(man *Man) Merried() string {
  man.Name = "Mr. "+man.Name
  return man.Name
}

func PointerMethod() {
  eko := Man{"Sunil"}
  eko.Merried()

  fmt.Println(eko)
}
