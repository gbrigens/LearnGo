package main

import "fmt"

func main() {
	fmt.Println("Array")

	var fruitList [4]string // You need to specify the number like as specified 4 or so

	fruitList[0] = "Apple"
	fruitList[1] = "Kiwi"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list: ", fruitList)      // It will show a blank for the missing  fruitList[2]
	fmt.Println("Fruit list: ", len(fruitList)) // It will show 4 but in real sense you have 3 values

	var vegList = [3]string{"Cabbage", "Wiki", "Onion"}
	fmt.Println("Vegetables list is: ", len(vegList))
}
