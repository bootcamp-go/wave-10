package main

import (
	"fmt"
	"time"
)

// MathOperation is a function type that takes two float64s and returns a float64
type MathOperation func(float64, float64) float64

// ShowTime is a function that takes a function type and two float64s and returns a float64
func ShowTime(mo MathOperation) MathOperation {
	return func(f1, f2 float64) float64 {
		// before
		fmt.Println(time.Now())
		// now
		return mo(f1, f2)
	}
}

// IncrementBy1 is a wrapper function that takes a function type and returns a function type
func IncrementBy1(mo MathOperation) MathOperation {
	return func(f1, f2 float64) float64 {
		// now
		result := mo(f1, f2)
		// after
		return result + 1
	}
}

// Add is a function that adds two float64s
func Add(a, b float64) float64 {
	return a + b
}

func main() {
	var mo MathOperation = ShowTime(Add)
	fmt.Println(mo(1, 2))

	var mo2 MathOperation = ShowTime(IncrementBy1(Add))
	fmt.Println(mo2(1, 2))
}