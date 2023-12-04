package main

import (
	"app/example-02/calculator"
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
		var mathError *calculator.MathError
		if errors.As(err, &mathError) {
			switch mathError.Code {
			case 1:
				fmt.Println("indeterminated math operation")
			case 2:
				fmt.Println("cannot divide by zero, invalid math operation")
			default:
				fmt.Println("an error occurred during the math operation")
			}
			return
		}
		fmt.Println("unknown error")
		return
	}

	// increment result
	result++
	fmt.Println(result)
}