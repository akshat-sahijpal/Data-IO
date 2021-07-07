package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ch0 := make(chan string)
	a := make(chan int, 100)
	b := make(chan int, 100)
	go worker(a, b)
	for i := 0; i < 100; i++ {
		a <- i
	}
	close(a)
	for i := 0; i < 100; i++ {
		fmt.Println(<-b)
	}

	go func() {
		for {
			ch <- "A" // Sender
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			ch0 <- "B" // Sender
			time.Sleep(time.Millisecond * 5000)
		}
	}()
	for {
		select { // Selects the channel thats available at that time
		case <-ch:
			fmt.Println(<-ch) // Receiver
		case <-ch0:
			fmt.Println(<-ch0) // Receiver
		}
	}
}
func workerThread(
	jobs <-chan int, //Only Receive
	results chan<- int /*Only Send */) {
}
func worker(jobs <-chan int /*Receive*/, res chan<- int /*Send*/) {
	for i := range jobs {
		res <- calc(i)
	}
}
func calc(i int) int {
	if i == 1 {
		return 1
	}
	return calc(i+1) + calc(i+2)
}
