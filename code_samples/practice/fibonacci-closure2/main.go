package main

import "fmt"

func fibo() func() int {
	a, b := 0, 1

	return func() int {
		fmt.Println("a: ", a, "b: ", b)
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
