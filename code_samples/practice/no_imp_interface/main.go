package main

import "fmt"

type pokemon struct {
	name string
	no int
}

func main()  {
	p := pokemon{
		name: "pikachu",
		no: 25,
	}
	fmt.Println(p) // {pikachu 25}
}
