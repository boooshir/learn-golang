package library

import (
	"container/list"
	"fmt"
)

func List() {
	var data *list.List = list.New()
	data.PushBack("eko")
	data.PushBack("kurniawan")
	data.PushBack("khenedy")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)
	next := head.Next()
	fmt.Println(next.Value)
	next = next.Next()
	fmt.Println(next.Value)

  // loop
  for e:= data.Front(); e!=nil; e = e.Next() {
    fmt.Println(e.Value)
  }
}
