package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	x := sum("james", xi...)
	fmt.Println("The total is", x)
}

// can only use ... in last argument
func sum(s string, x ...int) int {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	fmt.Println("The string is", s)
	return sum
}
