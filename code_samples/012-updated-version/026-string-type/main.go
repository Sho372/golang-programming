package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	s = "Hello, Hawaii"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// bite slice
	bs := []byte(s)
	fmt.Println(bs)        // UTF8(ASCII)
	fmt.Printf("%T\n", bs) //[]uint8 = []byte

	// byte loop
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) // %U: Unicode format
	}

	fmt.Printf("\n")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#x ", s[i]) // %#x: Hex format
	}
}
