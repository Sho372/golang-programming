package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	//x := func() int {
	//	return 451
	//}

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
	//clean up #1
	fmt.Println(x())
	//clean up #2
	fmt.Println(bar()())
}

func foo() string {
	return "Hello world"
	//return s
}

func bar() func() int {
	//anonymous func returning int
	return func() int {
		return 451
	}
}
