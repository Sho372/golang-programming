package main

import (
	"fmt"
	"math"
	"runtime"
)

var x int
var y float64

func main() {
	// Int8: Max positive number
	fmt.Println(math.Pow(2, 8)/2 - 1*+1)
	// Int8: Min negative number
	fmt.Println(math.Pow(2, 8) / 2 * -1)
	// UInt8: Max number (All positive)
	fmt.Println(math.Pow(2, 8) - 1*+1)

	x := 42
	y := 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
