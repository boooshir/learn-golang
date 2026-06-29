package deep

import (
	"fmt"
	"time"
)

type Cat struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

func DeepRun() {
	somethings := make(chan *Cat, 5)

	go func() {
		cat := &Cat{}
		cat.Name = "tompok"
		cat.Breed = " kucing kampung"
		somethings <- cat

	}()
	cat := <-somethings
	fmt.Println("cat name : ", cat.Name)
	fmt.Println("cat breed : ", cat.Breed)

}

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func RunJobs() {
	jobs := make(chan int, 2)
	results := make(chan int, 100)

	// start 3 worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	// collect results
	for a := 1; a <= 5; a++ {
		<-results
	}
}
