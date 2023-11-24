package main

import "fmt"

func main() {
	fmt.Println("Struct")

	// No inheritance in Go, no super or parent or child

	ger := User{"Ge", "ge@mail.com", true, 10}

	fmt.Println(ger)
	fmt.Printf("Usee details are: %+v\n", ger)
	fmt.Printf("Usee details are: %+v\n", ger.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
