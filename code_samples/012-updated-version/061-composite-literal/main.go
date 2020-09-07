package main

import "fmt"

func main() {
	// x := type{values} // composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	y := [10]int{}
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	// a SLICE allows you to group together VALUES of the same TYPE
}
