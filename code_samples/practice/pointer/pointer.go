package main

import "fmt"

var a = 100
var b = 200

func main() {
	//p := &a + 1; // no pointer arithmetic
	q := &b
	//fmt.Println(p)
	fmt.Println(q)
}
