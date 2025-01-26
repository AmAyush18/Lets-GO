package main

import "fmt"

// "defer" defers in LIFO
func main() {
	// defer fmt.Println("World!") // 4th
	// defer fmt.Println("One")    // 3rd
	// defer fmt.Println("Two")    // 2nd
	// fmt.Println("Hello")        // 1st

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
