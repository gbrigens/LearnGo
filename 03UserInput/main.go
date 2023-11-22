package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to simple app"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for Pizza: ") // No need for \ln

	// comma ok || err ok (Go does not have try catch) this is the use of comma or
	input, _ := reader.ReadString('\n') // _ for error in the case you catching
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T, ", input)

}

// Package resource link - pkg.go.dev

// Memmory
// New() Zero storage
// Make No Zero storage
