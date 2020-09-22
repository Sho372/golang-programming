package main

import (
	"flag"
	"fmt"
	"time"
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

	fmt.Println("s:", s)
	if *s != "" {
		fmt.Println("*s", *s)
		setDateStr = *s
		setDate = true
	}

	if setDate {
		parseDatetime(setDateStr)
	}
}

func parseDatetime(dateStr string) bool {
	// TODO: define custom layout
	yourDate, _ := time.Parse(time.Kitchen, dateStr)
	fmt.Println(yourDate)
	return true
}

func showDate() {

}