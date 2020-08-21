package main

import "fmt"

const (
	_  = iota             //0
	KB = 1 << (iota * 10) // 1 * 10
	MB = 1 << (iota * 10) // 2 * 10
	GB = 1 << (iota * 10) // 3 * 10
	TB = 1 << (iota * 10) // 4 * 10
)

func main() {
	x := 6
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)
	y = y << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Println()

	kb := 1024
	mb := 1024 * 1024
	gb := 1024 * 1024 * 1024
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	fmt.Println()

	fmt.Println("binary\t\t\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
