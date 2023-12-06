package main

import "fmt"

type Location struct {
	Latitude  float64
	Longitude float64
}

type Person struct {
	Name     string
	Age      int
	Location *Location
}

func main() {
	lc := Location{
		Latitude:  10.0,
		Longitude: 20.0,
	}
	p1 := Person{
		Name:     "John",
		Age:      20,
		Location: &lc,
	}
	p2 := p1

	p2.Name = "Jane"
	p2.Age = 30
	p2.Location.Latitude = 30.0
	p2.Location.Longitude = 40.0

	fmt.Printf("p1: name - %s, age - %d, location - %+v\n", p1.Name, p1.Age, p1.Location)
}
