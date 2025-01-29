package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to Post Form Request in Golang")

	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postForm"

	// formData
	data := url.Values{}
	data.Add("firstName", "Ayush")
	data.Add("lastName", "Dwivedi")
	data.Add("email", "amayush0818@gmail.com")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
