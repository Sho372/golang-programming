package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	//No error happens
	delete(m, "Ian Fleming")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Ian Fleming"])

	if _, ok := m["Miss Moneypenny"]; ok {
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)

}
