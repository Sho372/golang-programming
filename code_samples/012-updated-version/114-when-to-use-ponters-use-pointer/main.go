package main

import "fmt"

func main()  {
	x := 0
	foo(&x)
	fmt.Println(x)
}

func foo(y *int)  {
	fmt.Println(*y)
	*y = 42
	fmt.Println(*y)
}
