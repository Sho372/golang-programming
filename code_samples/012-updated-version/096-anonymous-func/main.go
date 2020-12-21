package main

import "fmt"

func main() {
	foo()

	// Anonymous func
	func() {
		fmt.Println("Anonymous func ran")
	}()

	// Anonymous func with parameter
	func(x int) {
		fmt.Println("The meaning of life", x)
	}(42)

}

func foo() {
	fmt.Println("foo ran")
}
