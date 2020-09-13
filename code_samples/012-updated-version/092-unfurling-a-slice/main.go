package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	//x := sum() // can pass zero value.
	fmt.Printf("%T\n", x)

	//fmt.Println(xi)
	//fmt.Println(&xi)
	//fmt.Printf("%T\n", xi)
	//fmt.Printf("%T\n", &xi)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += i
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}
