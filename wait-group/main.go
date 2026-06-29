package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter when the goroutine finished
	fmt.Printf("Worker %d starting\n", id)
	// simulate some work
	// ..
	time.Sleep(time.Second) // simulated work by sleeping
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment the counter for each goroutine
		go worker(i, &wg)
	}

	// launch an aditional goroutine after a delay
	time.Sleep(2 * time.Second)
	wg.Add(1)
	go worker(6, &wg)

	wg.Wait() // wait for all goroutine to finished
	fmt.Println("All worker finished.")
}
