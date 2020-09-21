package main

import (
	"flag"
	"fmt"
)

func usage(status int)  {
	fmt.Println("TODO: show usage")
}

func main()  {
	setDateStr := ""
	setDate := false
	fmt.Println(setDateStr)
	fmt.Println(setDate)

	// TODO: definition of flags
	s := flag.String("s", "", "set time described by STRING")
	flag.Parse()

	if s != nil {
		fmt.Println(*s)
		setDateStr = *s
		setDate = true
	}
}

func parseDatetime()  {
	
}

func showDate() {

}