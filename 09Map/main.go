package main

import "fmt"

func main() {
	fmt.Println("Map")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("List of languages: ", languages)
	fmt.Println("List of languages: ", languages["rb"])

	//Delate some values
	delete(languages, "rb")
	fmt.Println("List of languages: ", languages)

	// Loops
	for key, val := range languages {
		fmt.Println("For key %v, value is %v\n", key, val)
	}

}
