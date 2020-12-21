package main

import "fmt"

func main() {
	u := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"MoneyPenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(u.first)
	fmt.Println(u.friends)
	fmt.Println(u.favDrinks)

	for k, v := range u.favDrinks {
		fmt.Println(k, v)
	}
}
