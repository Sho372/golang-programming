package main

import "fmt"

func main()  {
	fmt.Println("2 + 3 =", mySum(2,3))
	fmt.Println("4 + 7 =", mySum(4,7))
	fmt.Println("10 + 100 =", mySum(10,100))
}

func mySum(s ...int) int {
	sum := 0
	for _ , v := range s {
		sum += v
	}
	return sum
}

