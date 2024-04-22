package main

import "fmt"

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Doubles(inputChannel <-chan int, outChannel chan<- int) {

	for value := range inputChannel {
		outChannel <- 2 * value
	}

	close(outChannel)
}

func PrintChannel(c <-chan int) {

	for value := range c {
		fmt.Println(value)
	}
}

func WRChannels() {
	inputChannel := make(chan int)
	outChannel := make(chan int)

	go Generator(inputChannel)
	go Doubles(inputChannel, outChannel)

	PrintChannel(outChannel)
}
