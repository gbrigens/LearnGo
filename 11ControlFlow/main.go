package main

import "fmt"

func main() {
	fmt.Println("Control Flow")

	loginCount := 20

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Pro User"
	} else {
		result = "No much"
	}

	fmt.Println(result)

	// CREATING ON the go

	if 9%2 == 0 {
		fmt.Println("Number even")
	} else {
		fmt.Println("Number odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is Not than 10")
	}
}
