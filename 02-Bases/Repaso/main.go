package main

import (
	"app/calculator"
	"errors"
	"os"
)

func main() {
	// calculator
	cl := calculator.NewDefault(10)

	// numbers
	n1 := 10.0
	n2 := 0.0

	// divide
	result, err := cl.Divide(n1, n2)
	if err != nil {
		// switch err {
		// 	case calculator.ErrDivisionByZero:
		// 		os.Exit(1)
		// 	case calculator.ErrDivisionIndeterminantForm:
		// 		os.Exit(2)
		// 	case calculator.ErrNoMoreOperationsAvailable:
		// 		os.Exit(3)
		// 	default:
		// 		os.Exit(127)
		// }
		switch {
			case errors.Is(err, calculator.ErrDivisionByZero):
				os.Exit(1)
			case errors.Is(err, calculator.ErrDivisionIndeterminantForm):
				os.Exit(2)
			case errors.Is(err, calculator.ErrNoMoreOperationsAvailable):
				os.Exit(3)
			default:
				os.Exit(127)
		}
	}

	// print result
	println(result)
}