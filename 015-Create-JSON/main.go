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
	EncodeJson()
}

func EncodeJson() {
	ctCourses := []course{
		{"Typescript Bootcamp", 59, "CodeTherapist.in", "Abc@123", []string{"web-dev", "ts"}},
		{"Javascript Bootcamp", 99, "CodeTherapist.in", "Def@123", []string{"web-dev", "js"}},
		{"React Bootcamp", 79, "CodeTherapist.in", "Ghi@123", nil},
	}

	// package this data as JSON data
	// finalJson, err := json.Marshal(ctCourses)
	finalJson, err := json.MarshalIndent(ctCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
