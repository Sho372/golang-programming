package main

import "fmt"

const (
	a = 2020 + iota // 0
	b = 2020 + iota // 1
	c = 2020 + iota // 2
	d = 2020 + iota // 3
)

func main() {
	fmt.Println(a, b, c, d)
}
