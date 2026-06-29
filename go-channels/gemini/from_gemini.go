package gemini

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs { // receive from jobs until it's close
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second) // simulate work
		result <- j * 2         // send the result
	}
}

func FromGemini() {
	const numbJobs = 5
	jobs := make(chan int, numbJobs)
	results := make(chan int, numbJobs)

	// start worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send jobs
	for j := 1; j <= numbJobs; j++ {
		jobs <- j
	}
	close(jobs) // signal that not more jobs will be sent

	//receive results
	for a := 1; a <= numbJobs; a++ {
		result := <-results
		fmt.Println("result :", result)
	}
}
