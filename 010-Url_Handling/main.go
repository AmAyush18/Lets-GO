package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://codetherapist.edu:8000/dev?course=js&paymentid=asdgv46687d"

func main() {
	fmt.Println("Handling URLS in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["course"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "codetherapist.edu",
		Path:    "/dev",
		RawPath: "user=ayush",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
