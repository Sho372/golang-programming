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


}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}
