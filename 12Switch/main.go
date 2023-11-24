package main

import (
	"math/rand"
	"time"
)

func main() {
	println("Switch statement")

	// Simple game
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	println("Value of the dice is ", diceNumber)

	switch diceNumber {
	case 1:
		println("Dice values is 1 you can open")
	case 2:
		println("You can move to 2 spot")
	case 3:
		println("You can move to 3 spot")
		fallthrough
	case 4:
		println("You can move to 4 spot")
	case 5:
		println("You can move to 5 spot")
	case 6:
		println("You can move to 6 spot")
	default:
		println("Oops!!")
	}

}
