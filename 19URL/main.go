package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=frdexfbhjgcde"

func main() {
	println("URL Handling")

	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// Storing in var
	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Params are ", val)
	}

	// Construction of URl

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/tutcss",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
