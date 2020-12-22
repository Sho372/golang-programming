package main

import "fmt"

func main() {
	i := 1986
	for {
		fmt.Println(i)
		i++
		if i > 2021 {
			break
		}
	}

}
