package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Printf("Worker %d waiting for jobs\n", id)
	for j := range jobs {
		fmt.Println("worker", id, "started job")
		time.Sleep(time.Second)
		fmt.Println("worker", id, " finished jon")
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	time.Sleep(time.Second * 3)

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
