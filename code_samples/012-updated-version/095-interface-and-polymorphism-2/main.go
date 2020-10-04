package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// implementation of human interface
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the secretAgent speak")
}

// implementation of human interface
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)

	}
}

func main() {
	// this value's type is secretAgent and human in this time.
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
