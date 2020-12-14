package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))

	nowUTC := now.UTC()
	fmt.Println(nowUTC.Format(time.RFC3339))

	jst := time.FixedZone("Asia/Tokyo", int((9 * time.Hour).Seconds()))
	pdt := time.FixedZone("PDT", int((-7 * time.Hour).Seconds()))
	dateJst := time.Date(2020, 9, 11, 3, 0, 0, 0, jst)
	fmt.Println(dateJst.Format(time.RFC3339))
	dateUTC := dateJst.UTC()
	fmt.Println(dateUTC.Format(time.RFC3339))

	datePdt := time.Date(2020, 9, 11, 3, 0, 0, 0, pdt)
	fmt.Println(datePdt.Format(time.RFC3339))
	dateUTC2 := datePdt.UTC()
	fmt.Println(dateUTC2.Format(time.RFC3339))

}
