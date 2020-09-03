package main

import "fmt"

func main()  {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &a)

	fmt.Println(*b) // dereference
	*b = 42

	fmt.Println(a)
}
