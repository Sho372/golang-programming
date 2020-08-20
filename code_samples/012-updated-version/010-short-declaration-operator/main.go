package main

import "fmt"

func main() {
	x := 42
	fmt.Print(x)
	x = 99
	fmt.Println(x)
	y := 100 + 70
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
