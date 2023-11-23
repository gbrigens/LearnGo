package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Orange", "Apple", "Grape"}

	fmt.Println("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // Slicing [1:] - This removes the first element from fruit list and returns all remaining elements in another slice with updated length

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // Slicing [1:3] - It will include the 1,2 only. As it's a range of 2 elements, the length is reduced to 2, so only two fruits are printed.
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 243
	highScore[1] = 244
	highScore[2] = 245
	highScore[3] = 246
	// highScore[4] = 246 // wont print out of bounds error - it will only print up to the length

	highScore = append(highScore, 555, 444, 321) // It will include this as it will increase the slice len

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// Performing other operations
	sort.Ints(highScore)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// Removing values from Slice based on index

	var courses = []string{"React", "Swift", "Python", "Go"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // It will remove the python from index 2 and other will remain...

	fmt.Println(courses)
}
