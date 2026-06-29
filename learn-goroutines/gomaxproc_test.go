package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxproc(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}
	totalCPU := runtime.NumCPU()
	fmt.Println("total cpu :", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
	group.Wait()
}
