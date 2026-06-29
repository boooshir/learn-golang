package library

import (
	"errors"
	"fmt"
)

var (
  ValidationError = errors.New("ValidationError")
  NotFoundError = errors.New("not found error")
)

func getById(id string) error {
  if id == "" {
    return ValidationError
  } 

  if id != "eko" {
    return NotFoundError
  }
  return nil
}

func Error() {
  err := getById("")

  if err != nil {
    if errors.Is(err, ValidationError) {
      fmt.Println("validation error")
    } else if errors.Is(err, NotFoundError) {
      fmt.Println("notfound error")
    }
  } else {
    fmt.Println("success")
  }
}
