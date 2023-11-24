package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	println("Web Request")

	response, err := http.Get(url)

	checkNillErr(err)

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // Callers responsibility to close the connection.

	//Reading
	databytes, err := io.ReadAll(response.Body)
	checkNillErr(err)

	content := string(databytes)

	fmt.Printf(content)

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
