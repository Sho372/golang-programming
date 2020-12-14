package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a flag - name, default value, help message.
	nFlag := flag.Int("n", 1234, "help message for flag n")
	flag.Parse()

	// nFlag refers the value of flag
	fmt.Println(*nFlag)
	if *nFlag == 1234 {
		flag.PrintDefaults()
	}
}
