package main

import "fmt"

func main() {
	println("Functions")
	greeter() // Calling the function

	result := adder(3, 5)

	fmt.Println("Result is:", result)

	proResult := prAdder(2, 3, 4, 5, 6, 7, 8)

	fmt.Println("Pro result is:", proResult)
}

func greeter() {
	println("Jambo")
}

// Variadic functions
func prAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
