package library

import (
	"fmt"
	"time"
)

func Duration() {
  var duration1 time.Duration = 100 * time.Second
  var duration2 time.Duration = 10 * time.Microsecond
  var duration3 time.Duration = duration1 - duration2

  fmt.Println(duration3)
  fmt.Printf("duration: %d \n", duration3)
}
