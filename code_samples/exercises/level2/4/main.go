package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)
	xx := x << 1
	fmt.Printf("%d\t%b\t%#X", xx, xx, xx)
}
