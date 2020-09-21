package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main()  {
	// Subcommands
	showCommand := flag.NewFlagSet("show", flag.ExitOnError)

	// show subcommand flag pointers
	//utc := showCommand.String("u", "UTC", "UTC")
	//jst := showCommand.String("j", "Asia/Tokyo", "Asia/Tokyo")

	if len(os.Args) < 2 {
		fmt.Println("show is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "show":
		showCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	now := time.Now()
	zone, offset := now.Zone()
	println(zone, offset)
	nowUTC := now.UTC()

	jst := time.FixedZone("Asia/Tokyo", int((9*time.Hour).Seconds()))
	pdt := time.FixedZone("PDT", int((-7*time.Hour).Seconds()))

	if showCommand.Parsed() {
		if len(os.Args) == 2 {
			nowPdt := nowUTC.In(pdt)
			fmt.Println(nowPdt.Format(time.RFC3339))
			fmt.Println(nowUTC.Format(time.RFC3339))
			nowJST := nowUTC.In(jst)
			fmt.Println(nowJST.Format(time.RFC3339))
		}
		if len(os.Args) == 3 {
			layout := "2006-01-02-15-04"
			str := os.Args[2]
			t, err := time.Parse(layout, str)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(t)
		}
	}

}
