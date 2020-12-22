package main

import "fmt"

func main() {
	fmt.Println(4*3*2*1)
	fmt.Println(factrio(4))
	fmt.Println(factrioLoop(4))

}

func factrio(n int) int {
	if n == 1 {
		return 1
	}
	return n * factrio(n-1)
}

func factrioLoop(n int) int {
	ret := 1
	for i := 1; i <= n; i++ {
		ret *= i
	} 
	return ret
}
