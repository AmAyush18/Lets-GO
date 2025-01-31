package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // backticks (``) are for giving kind of aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // - -> entirely removes the field
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Let's Create some JSON")
	DecodeJson()
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Typescript Bootcamp",
		"Price": 59,
		"website": "CodeTherapist.in",
		"tags": ["web-dev","ts"]
	}
	`)

	var ctCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &ctCourse)
		fmt.Printf("%#v", ctCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v", ctCourse)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value %v and Typeis :%T\n", k, v, v)
	}
}
