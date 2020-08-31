package main

import "fmt"

// we create VALUES of a certain TYPE that are stored in VARIABLES and those VARIABLES have IDENTIFIERS

var x int

type person struct {
	first string
	last string
}

type foo int
var y foo
var z foo
const bar1 = 42
const bar2 int = 42 // Here is no longer a constant of a kind.

func main()  {

	//y = int(42)
	y = 42 // This 42 is indeterminate. There is an indeterminate type in Go where it's like something constants can be a constant of a kind. So we can convert to int.
	fmt.Println(y)
	fmt.Printf("%T\n", y) // type is main.foo, not int
	fmt.Printf("%T\n", int(y)) // type is int. can convert foo to int

	z = bar1 // z: foo, bar1: indeterminate -> can be converted
	fmt.Printf("%T\n", z)
	fmt.Println(z)

	//z = bar2 // z: foo, bar1: int -> can not be converted
	fmt.Printf("%T\n", z)
	fmt.Println(z)

}
