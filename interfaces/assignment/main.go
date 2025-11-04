package main

import "fmt"

type square struct {
	name       string
	sideLength float64
}

type triangle struct {
	name   string
	height float64
	base   float64
}

type shape interface {
	getName() string
	getArea() float64
}

func main() {
	square := square{name: "Square", sideLength: 10}
	triangle := triangle{name: "Triangle", height: 10, base: 3}

	printArea(square)
	printArea(triangle)
}

// functions that recieve an interface are NOT reciever functions
// In Go, you cannot define methods on interface types directly. Methods can only be defined on concrete types
func printArea(s shape) {
	fmt.Println("Area of", s.getName(), "is", s.getArea())
}

func (t triangle) getName() string {
	return t.name
}

func (s square) getName() string {
	return s.name
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
