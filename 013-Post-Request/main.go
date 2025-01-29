package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Post Json Request in Golang")

	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's Go",
			"price": 0,
			"platform": "Code Therapist"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
