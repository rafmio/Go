// https://habr.com/ru/articles/852556/
/*
WorkerPool is great for situations where you need to process a large number of similar tasks,
for example, incoming API requests, working with files, database requests and other tasks
that require parallel processing. The pattern allows you to effectively distribute the load
without creating an excessive number of goroutines.
*/
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Duration(job) * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		result <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// creating 3 workers pool
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send tasks to channel jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// close the task channel so that the workers understand that there will be no more tasks
	close(jobs) //

	// receive results from channel results
	for r := 1; r <= numJobs; r++ {
		res := <-results
		fmt.Printf("Result: %d\n", res)
	}

	close(results)
}
