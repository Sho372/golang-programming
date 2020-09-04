package main

import "fmt"

// no pointer
func main()  {
	x := 0
	foo(x)
	fmt.Println(x)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}
