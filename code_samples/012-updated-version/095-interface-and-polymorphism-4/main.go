package main

import (
	"fmt"

	"./myFmt"
)

type Pokemon struct {
	Name  string
	Level uint
}

// String makes Pokmeon satisfy the myFmt.Stringer interface.
func (p Pokemon) String() string {
	return fmt.Sprintf("%v (%d)", p.Name, p.Level)
}

func (p Pokemon) Move() {
	fmt.Println("Electro Ball")
}

type Trainer struct {
	Name  string
	Badge int
}

// String makes Trainer satisfy the myFmt.Stringer interface.
func (p Trainer) String() string {
	return fmt.Sprintf("%v (%d)", p.Name, p.Badge)
}

func (p Trainer) Speak() {
	fmt.Println("Gotcha")
}

type Item struct {
	Name   string
	Number int
}

// String makes Item satisfy the myFmt.Stringer interface.
func (i Item) String() string {
	return fmt.Sprintf("%v (%d)", i.Name, i.Number)
}

func (i Item) Use() {
	fmt.Println("Great")
}

func getString(v myFmt.Stringer) {
	switch i := v.(type) {
	case Pokemon:
		i.Move()
	case Trainer:
		i.Speak()
	case Item:
		i.Use()
	}
}

func main() {

	p := Pokemon{
		Name:  "Pikachu",
		Level: 50,
	}

	t := Trainer{
		Name:  "Ash",
		Badge: 8,
	}

	i := Item{
		Name:   "Potion",
		Number: 20,
	}

	getString(p)
	getString(t)
	getString(i)

}
