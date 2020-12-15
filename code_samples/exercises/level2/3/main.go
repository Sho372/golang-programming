package main

import "fmt"

func main() {
	const (
		x        = "hello" // untyped constant
		y string = "world" // typed constant
	)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(x, y)
}
