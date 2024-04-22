package main

import "fmt"

func BufferedExample() {
	c := make(chan int, 1) // buffered channel

	c <- 1

	fmt.Println(<-c)
}
