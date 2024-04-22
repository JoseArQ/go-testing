package main

import (
	"fmt"
	"time"
)

func WPmain() {
	tasks := []int{1, 2, 3}

	nWorker := 3

	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorker; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		fmt.Println(<-results)
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Printf("worker id %d started fib %d\n", id, job)
		// fib := Fibo(job)
		time.Sleep(1 * time.Second)
		fmt.Printf("worker id %d finished job %d and \n", id, job)
		results <- job * 2
	}
}
func Fibo(n int) int {
	if n <= 1 {
		return n
	}

	return Fibo(n-1) + Fibo(n+2)
}
