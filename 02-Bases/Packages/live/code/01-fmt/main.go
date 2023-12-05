package main

import "fmt"

func main() {
	// fmt cheat sheet
	// - %v: value in default format
	// - %d: integer
	// - %f: float
	// - %s: string
	// - %t: boolean
	// - %p: pointer
	//
	// - %T: type of the value
	// - %b: binary
	// - %c: character
	// - %x: hexadecimal
	fmt.Printf("- struct: %v\n- integer: %d\n- float: %f\n- string: %s\n- boolean: %t\n- pointer: %p\n- type: %T\n",
		struct{ Name string }{Name: "John"}, 1, 1.1, "hello", true, new(int), struct{ Name string }{Name: "John"},
	)
}
