package main

import "fmt"

func main() {
	println("Loops")

	days := []string{"sunday", "Tuesday", "Wednesday", "Thursdays", "Friday", "Saturday"}
	fmt.Println(days)

	// Run through
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	// or

	// for d := range days {
	// 	fmt.Println(d)
	// 	fmt.Println(days[d])
	// }

	// or
	// for index, day := range days {
	// 	fmt.Println(index, day)
	// }

	// While loop alternative in go
	myValue := 1

	for myValue < 10 {

		if myValue == 2 {
			goto gotome
		}
		if myValue == 5 {
			// break //stop

			// Skip specified value in this case 5
			myValue++
			continue

		}
		fmt.Println("Value is: ", myValue)
		myValue++
	}

	// Provide go to statement
gotome:
	println("Just jump to this line of code")

}
