package main

import "fmt"

func main() {

	fmt.Println("Pointers")
	// Absolute guarantee that the value is passed by asses the address.

	var ptr *int

	fmt.Println("Value of pointer is ", ptr)

	myNum := 20

	// Var that pint to a address
	var myPtr = &myNum // & used to reference

	fmt.Println("Value of actual pointer is ", myPtr)
	fmt.Println("Value of actual pointer is ", *myPtr)

	*myPtr = *myPtr * 2
	fmt.Println("New value is: ", myNum)
}
