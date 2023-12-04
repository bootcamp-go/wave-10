package main

import (
	"app/example-00/calculator"
	"fmt"
)

func main() {
	// numbers
	n1 := 10.0
	n2 := 0.0

	// divide
	result, err := calculator.Div(n1, n2)
	if err != nil {
		switch err {
		case calculator.ErrDivisionByZero:
			fmt.Println("invalid operation: division by zero")
		default:
			fmt.Println(err)
		}
		return
	}

	// increment result
	result++
	fmt.Println(result)
}