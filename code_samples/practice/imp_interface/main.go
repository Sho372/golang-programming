package main

import "fmt"

type pokemon struct {
	Name string
	No int
}

// String makes pokemon satisfy the Stringer interface
func (p pokemon) String() string {
	return fmt.Sprintf("%v (%d)", p.Name, p.No)
}

func main()  {
	p := pokemon{
		Name: "pikachu",
		No: 25,
	}
	fmt.Println(p) // pikachu (25)
}

