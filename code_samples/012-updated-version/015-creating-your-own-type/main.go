package main

import (
	"fmt"
)

var a int

type hotdog int // hotdog is int
var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b) // ${package name}.${type name}

	// You can not
	// a = b
}
