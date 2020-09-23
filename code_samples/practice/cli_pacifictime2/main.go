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


func parseDatetime(setDateStr string) (bool, time.Time) {
	// TODO: define custom layout
	// 2020-09-23-21:00
	yourDate, _ := time.Parse("2006-01-02-15:04", setDateStr)
	fmt.Println(yourDate)
	return true, yourDate
}

func showDate() {

}