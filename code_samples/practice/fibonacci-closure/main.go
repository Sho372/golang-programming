package main

import "fmt"

func main() {
	f := fibo(13)
	fmt.Println(f()) //n:  5 a: 0  b: 1
	fmt.Println(f()) //n:  2 a: 2  b: 3
}

func fibo(n int) func() int {
	a, b := 0, 1

	return func() int {
		fmt.Println("n: ", n, "a:", a, " b:", b)
		if n < 3 {
			return 1
		}

		for n > 0 {
			a, b = b, a+b
			n--
		}
		return a
	}
}
