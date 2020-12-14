package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 50)
	c := (42 >= 50)
	d := (42 != 50)
	e := (42 < 50)
	f := (42 > 50)
	fmt.Println(a, b, c, d, e, f)
}
