package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

// implementation of human interface
func (s secretAgent) speak()  {
	fmt.Println("I am", s.first, s.last)
}

// implementation of human interface
func (p person) speak()  {
	fmt.Println("I am", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human)  {
	fmt.Println("I was passed into bar", h)
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

	p1 := person{
		first: "Dr.",
		last: "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)
}