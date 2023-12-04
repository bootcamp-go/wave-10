package main

import (
	"app/example-01/calculator"
	"errors"
	"fmt"
)

func main() {
	// numbers
	n1 := 10.0
	n2 := 0.0

	// divide
	result, err := calculator.Div(n1, n2)
	if err != nil {
		if errors.Is(err, calculator.ErrDivisionByZero) {
			fmt.Println("invalid operation: division by zero")
			return
		}
		fmt.Println(err)
		return
	}

	// increment result
	result++
	fmt.Println(result)
}