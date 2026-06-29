package library

import (
	"fmt"
	"math"
)

func Math() {
  fmt.Println(math.Ceil(1.40))
  fmt.Println(math.Floor(1.60))
  fmt.Println(math.Round(1.60))
  fmt.Println(math.Max(10, 50))
  fmt.Println(math.Min(40, 30))
}
