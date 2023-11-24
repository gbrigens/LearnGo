package main

import "fmt"

func main() {
	println("Defer Statements in Go!")

	defer fmt.Println("Hey hey")
	defer fmt.Println("Hey Two")
	defer fmt.Println("Hey Three") //3 and so on upwards
	fmt.Println("Krusty here")     // 1

	myDefer() // 2

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
