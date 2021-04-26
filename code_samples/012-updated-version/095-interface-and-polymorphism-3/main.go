package main

import "fmt"

func main() {

	var i interface{} = "Hello, world!"

	// type assertions
	s, ok := i.(string)
	fmt.Println(s, ok)
	fmt.Printf("%T\n", s) // s is string type

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// panic
	f = i.(float64)
	fmt.Println(f)
}
