package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	LAYOUT = "2006-01-02-15:04MST"
	PDT = "PDT"
	PST = "PST"
	JST = "JST"
	UTC = "UTC"
)

func usage(status int) {
	//TODO: Prepare a usage for this command
	fmt.Println("TODO: show usage")
}

func main() {
	setDateStr := ""
	setDate := false
	setTime := time.Now()
	setTimeZoneStr, _ := setTime.Zone()

	s := flag.String("s", "", "set time described by STRING")
	z := flag.String("z", "", "set time zone described by STRING")
	flag.Parse()

	switch *z {
	case strings.ToLower(PDT):
		setTimeZoneStr = PDT
		pdtZone := time.FixedZone(PDT, -7*60*60)
		setTime = time.Date(setTime.Year(), setTime.Month(), setTime.Day(), setTime.Hour(), setTime.Minute(), setTime.Second(), setTime.Nanosecond(), pdtZone)

	case strings.ToLower(JST):
		setTimeZoneStr = JST
	case strings.ToLower(UTC):
		setTimeZoneStr = UTC
	}

	if *s != "" {
		setDateStr = *s
		setDate = true
	}

	if setDate {
		ok := true
		setTime, ok = parseDatetime(setDateStr, setTimeZoneStr)
		if !ok {
			os.Exit(1)
		}
	}
	showDate(setTime)
}

func parseDatetime(setDateStr string, setTimeZoneStr string) (time.Time, bool) {
	switch setTimeZoneStr {
	case PDT:
		{
			setDateStr = setDateStr + PDT
		}
	case JST:
		{
			setDateStr = setDateStr + JST
		}
	case UTC:
		{
			setDateStr = setDateStr + UTC
		}
	}

	// Layout: 2020-09-23-21:00
	ok := true
	t, err := time.Parse(LAYOUT, setDateStr)
	if err != nil {
		fmt.Println(err)
		ok = false
	}
	return t, ok
}

func showDate(setTime time.Time) {
	timeZone, _ := setTime.Zone()
	switch timeZone {
	case PDT:
		{
			fmt.Println()

			// setTimeZone
			showFormattedDate(setTime, timeZone)
			fmt.Println("----------------------")

			// UTC
			utcZone, err := time.LoadLocation(UTC)
			if err != nil {
				fmt.Println(err)
			}
			utcTime := setTime.In(utcZone)
			showFormattedDate(utcTime, utcZone.String())
		
			// JST
			jstZone := time.FixedZone(JST, 9*60*60)
			jstTime := setTime.In(jstZone)
			showFormattedDate(jstTime, jstZone.String())
		}
	case JST:
		{

			fmt.Println()

			// setTimeZone
			showFormattedDate(setTime, timeZone)
			fmt.Println("----------------------")

			// PDT
			pdtZone := time.FixedZone(PDT, -7*60*60)
			pdtTime := setTime.In(pdtZone)
			showFormattedDate(pdtTime, pdtZone.String())

			// UTC
			utcZone, err := time.LoadLocation(UTC)
			if err != nil {
				fmt.Println(err)
			}
			utcTime := setTime.In(utcZone)
			showFormattedDate(utcTime, utcZone.String())
		}
	case UTC:
		{
			fmt.Println()

			// setTimeZone
			showFormattedDate(setTime, timeZone)
			fmt.Println("----------------------")

			// PDT
			pdtZone := time.FixedZone(PDT, -7*60*60)
			pdtTime := setTime.In(pdtZone)
			showFormattedDate(pdtTime, pdtZone.String())

			// JST
			jstZone := time.FixedZone(JST, 9*60*60)
			jstTime := setTime.In(jstZone)
			showFormattedDate(jstTime, jstZone.String())
		}
	}
}

func showFormattedDate(setTime time.Time, timeZone string)  {
	fmt.Println(setTime.Format("[MST] 2006-01-02 15:04"))
}
