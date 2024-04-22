package main

import (
	"fmt"
	"sync"
	"time"
)

func Semaforos() {

	c := make(chan int, 3)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomethingTwo(i, &wg, c)
	}
	wg.Wait()
}

func doSomethingTwo(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("ID %d started\n", i)
	time.Sleep(3 * time.Second)
	fmt.Println("end")
	<-c
}
