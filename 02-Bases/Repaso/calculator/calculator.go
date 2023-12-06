package calculator

import (
	"errors"
	"fmt"
)

/*
	This package is used to perform basic mathematical operations
	like addition, subtraction, multiplication and division
*/

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrDivisionIndeterminantForm = errors.New("division resulted in indeterminant form")
	ErrNoMoreOperationsAvailable = errors.New("no more operations available")
)

func NewDefault(limit int) *Default {
	return &Default{
		Limit: limit,
		Count: 0,
	}
}

// Default is the default implementation of Calculator
type Default struct {
	// Limit is the maximum number of operations allowed
	Limit int
	// Count is the number of operations performed
	Count int
}

// Divide is a function that divides two numbers
func (d *Default) Divide(a, b float64) (result float64, err error) {
	// check limit of operations
	if d.Count >= d.Limit {
		err = fmt.Errorf("%w: %d", ErrNoMoreOperationsAvailable, d.Limit)
		return
	}

	// check indeterminant form
	if a == 0 && b == 0 {
		err = fmt.Errorf("%w: %f, %f", ErrDivisionIndeterminantForm, a, b)
		return
	}
	// check division by zero
	if b == 0 {
		err = fmt.Errorf("%w: %f, %f", ErrDivisionByZero, a, b)
		return
	}

	result = a / b

	// increment count
	d.Count++

	return
}