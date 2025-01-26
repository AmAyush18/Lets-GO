package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.github.com/users/AmAyush18"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// %T -> type of
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // Caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
