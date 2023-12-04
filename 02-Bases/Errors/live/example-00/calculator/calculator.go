package calculator

import (
	"errors"
)

var (
	// ErrDivisionByZero is returned when the divisor is zero
	ErrDivisionByZero = errors.New("division by zero")
)

// Sum returns the sum of two float64 numbers
func Sum(a, b float64) (result float64) {
	result = a + b
	return
}

// Sub returns the subtraction of two float64 numbers
func Sub(a, b float64) (result float64) {
	result = a - b
	return
}

// Mul returns the multiplication of two float64 numbers
func Mul(a, b float64) (result float64) {
	result = a * b
	return
}

// Div returns the division of two float64 numbers
func Div(a, b float64) (result float64, err error) {
	if b == 0 {
		err = ErrDivisionByZero
		return
	}
	result = a / b
	return
}