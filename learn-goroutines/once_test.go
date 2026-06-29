package main

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOne() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOne)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
