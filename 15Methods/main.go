package main

import "fmt"

func main() {
	println("Methods")

	ger := User{"Ge", "ge@mail.com", true, 10}

	fmt.Println(ger)
	fmt.Printf("Usee details are: %+v\n", ger)
	fmt.Printf("Usee details are: %+v\n", ger.Name)

	ger.GetStatus()
	ger.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Method : - To get user if is online or not
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// this will pass a copy and not original Email. To get original email you need a pointer
func (u User) NewMail() {

	u.Email = "me@mail.dev"
	fmt.Println("Email of the user is: ", u.Email)
}
