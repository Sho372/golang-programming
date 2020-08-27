package main

import "fmt"

func main() {

	x := make([]int, 10, 12) // length: 10, capacity: 12
	fmt.Println(x)
	fmt.Println(len(x)) // length: 10
	fmt.Println(cap(x)) // capacity: 12

	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x)) // length: 11
	fmt.Println(cap(x)) // capacity: 12

	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x)) // length: 12
	fmt.Println(cap(x)) // capacity: 12

	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x)) // length: 13
	fmt.Println(cap(x)) // capacity: 24

	// initialize slice
	s := make([]int, 5, 10)
	fmt.Println(s)
}
