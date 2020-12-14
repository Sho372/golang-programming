package main

import "fmt"

var x = 42;
var y = "James Bond"
var z = true;

func main() {
	// short declaration
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// multiple print statements
	fmt.Println(x, y, z)

	s := fmt.Sprintf("%s is %d years old. \n", y, x)
	fmt.Println(s)

}
