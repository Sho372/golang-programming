package main

import `fmt`

func main() {
	// raw string literals
	rs1 :=
`hello 
world`
	fmt.Println(rs1)
	// interpreted string literals
	is1 := "hello\nworld"
	fmt.Println(is1)

	rs2 := `"hello"`
	fmt.Println(rs2)

	is2 := "\"hello\""
	fmt.Println(is2)
	
	rs3 := `\a`
	fmt.Println(rs3)

	is3 := "\a"
	fmt.Println(is3)
}
