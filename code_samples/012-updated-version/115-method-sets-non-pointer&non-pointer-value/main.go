package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// NO POINTER RECEIVER
// it's not a pointer for the receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// interface typeのshapeを引数で受け取る。 circle typeではない
func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {

	//c := circle{
	//	radius: 5,
	//}

	// shorthand
	c := circle{5}
	info(c)
}
