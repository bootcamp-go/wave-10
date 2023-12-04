package calculator

import (
	"fmt"
)

// MathError is a custom error type
type MathError struct {
	Msg  string
	Op   string
	Code int
}

// Error returns the error message
func (e *MathError) Error() string {
	return e.Msg
}

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
	if a == 0 && b == 0 {
		err = &MathError{
			Msg:  fmt.Sprintf("indeterminated: %v/%v", a, b),
			Op:   "division",
			Code: 1,
		}
		return
	}
	if b == 0 {
		err = &MathError{
			Msg:  fmt.Sprintf("cannot divide by zero: %v/%v", a, b),
			Op:   "division",
			Code: 2,
		}
		return
	}
	result = a / b
	return
}