package main

import (
	"fmt"
	"time"
)

func mathOperation(a, b int) (result int) {
	fmt.Println("Math operation started")
	time.Sleep(1 * time.Second)
	result = a + b
	fmt.Println("Math operation ended")
	return
}

func main() {
	// start timer
	st := time.Now()

	// do some math operation
	go func() {
		fmt.Println("Result: ", mathOperation(3, 4))
	}()
	go func() {
		fmt.Println("Result: ", mathOperation(5, 6))
	}()
	go func() {
		fmt.Println("Result: ", mathOperation(7, 8))
	}()

	// wait for 2 seconds
	time.Sleep(1500 * time.Millisecond)

	// lapse
	duration := time.Since(st)
	fmt.Println("Duration: ", duration.Milliseconds())
}