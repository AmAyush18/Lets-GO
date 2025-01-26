package main

import "fmt"

func main() {
	fmt.Println("Structs and Methods in golang")
	// no inheritance in golang; No super or parent

	ayush := User{"Ayush", "amayush0818@gmail.com", true, 22}

	// fmt.Println(ayush)
	// fmt.Printf("ayush' details are: %+v\n", ayush)
	// fmt.Printf("Name is %v and email is %v.", ayush.Name, ayush.Email)

	ayush.GetStatus()
	ayush.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Current Status is", u.Status)
}

// we get a copy and not actual reference of the value for that we can use pointer
func (u User) NewMail() {
	u.Email = "hacker@golang.dev"
	fmt.Println("New Mail: ", u.Email)
}
