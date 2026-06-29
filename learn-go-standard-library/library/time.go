package library

import (
	"fmt"
	"time"
)

func Time() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 9, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(parse)
}
