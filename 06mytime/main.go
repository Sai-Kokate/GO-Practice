package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Go")

	myTime := time.Now()

	fmt.Println(myTime)

	formattedTime := myTime.Format("01-02-2006")

	fmt.Println(formattedTime)
	fmt.Println(myTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 15, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
