package library

import (
	"fmt"
	"reflect"
)

type Sample struct {
  Name string
}

func Reflection() {
  sample := Sample{"Eko"}
  sampleType := reflect.TypeOf(sample)
  fmt.Println(sampleType)
  structField := sampleType.Field(0)

  fmt.Println(structField.Name)
  fmt.Println(structField.Type)
}
