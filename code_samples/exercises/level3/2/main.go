package main

import "fmt"

func main() {
	var i rune = 65
	for ; i <= 90; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
