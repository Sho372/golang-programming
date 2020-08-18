package main

import "fmt"

func main() {
	fmt.Println("Hello everyone")

	foo()
	fmt.Println("something more")
	bar()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	printlnTest()

}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}

func printlnTest() {
	// Use the returns
	n, err := fmt.Println("Hello, world", 42, true)
	fmt.Println(n)
	fmt.Println(err)

	// Throw away one of returns with _
	n2, _ := fmt.Println("Hello, world", 42, true)
	fmt.Println(n2)

	// You can't throw away without _
	//n3, e := fmt.Println("Hello, world", 42, true)
	//fmt.Println(n3)

}
