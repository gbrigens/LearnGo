package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time, GO")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Local().Format("01-02-2006 15:01.09 Monday"))

	createdDate := time.Date(2020, time.December, 10, 24, 21, 0, 0, time.Local)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01/02/2005 Mon"))

}
