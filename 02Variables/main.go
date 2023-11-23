package main

import "fmt"

// Global
const LoginToken string = "fjdfeufuef" // Capital letter means public

func main() {
	var username string = "Ge"

	fmt.Printf("Variable is of type: %T \n", username)

	var isLogedin bool = false
	fmt.Println(isLogedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)

	var smallFloat float64 = 255.13354545
	fmt.Println(smallFloat)

	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type
	var website = "my fancy web"
	fmt.Println(website)

	// no var style
	numberOfUsers := 2300
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
