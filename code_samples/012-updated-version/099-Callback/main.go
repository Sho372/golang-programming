package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(ii...))

	fmt.Println(even(sum, ii...))

}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
