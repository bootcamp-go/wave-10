package switcher

import "fmt"

func A() {
	fmt.Println("A")
}

func B() {
	fmt.Println("B")
}

func C() {
	fmt.Println("C")
}

func Default() {
	fmt.Println("Default")
}

func Handler() func() int {
	return func() int {
		return 5
	}
}

func Switcher(name string) (fn func()) {
	switch name {
	case "A":
		fn = A
	case "B":
		fn = B
	case "C":
		fn = C
	default:
		fn = A
	}
	return
}