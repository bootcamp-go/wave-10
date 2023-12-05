package main

import (
	"fmt"
	"time"
)

func mathOperation(a, b int, ch chan<- int) {
	fmt.Println("Math operation started")
	time.Sleep(1 * time.Second)
	result := a + b

	// send result to channel
	ch <- result
	close(ch)

	fmt.Println("Math operation ended")
	return
}

func main() {
	// start timer
	st := time.Now()

	// make channel
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go mathOperation(3, 4, ch1)
	go mathOperation(5, 6, ch2)
	go mathOperation(7, 8, ch3)

	r1 := <-ch1
	r2 := <-ch2
	r3 := <-ch3
	fmt.Println("Result: ", r1, r2, r3)

	// lapse
	duration := time.Since(st)
	fmt.Println("Duration: ", duration.Milliseconds())
}