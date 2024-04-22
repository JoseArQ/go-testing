package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoDelay(d1, c1, 1)
	go DoDelay(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMessage1 := <-c1:
			fmt.Println(channelMessage1)
		case channelMessage2 := <-c2:
			fmt.Println(channelMessage2)
		}
	}
}

func DoDelay(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
