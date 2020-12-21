package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	v         vehicle
	fourWheel bool
}

type seadan struct {
	v      vehicle
	luxury bool
}

func main() {
	// composit literal
	v := vehicle{
		doors: 4,
		color: "red",
	}
	fmt.Println(v)

	// composit literal
	t := truck{
		v: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	s := seadan{
		v: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.fourWheel)
	fmt.Println(s.luxury)
}
