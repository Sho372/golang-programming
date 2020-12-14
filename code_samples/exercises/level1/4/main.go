package main

import "fmt"

//Type declaration
type weight int //underlying type is int

//Type declarations
type (
	age    int   //Type defnition. underlying type is int
	height = int //Alias declarations. underlying type is int

)

var x age

func main() {
	fmt.Printf("%T\n", x)
	fmt.Printf("My age is %d\n", x) // ZERO value -> 0
	x = 42
	fmt.Printf("My age is %d", x)
}
