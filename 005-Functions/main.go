package main

import "fmt"

func main() {
	fmt.Println("Let's write some functions in golang!")
	// anniversaryWish()
	// birthdayWish()

	result := sum(18, 8)
	fmt.Println("Sum of 18 and 8 is", result)

	proResult, proMessage := proSum(18, 8, 9, 10, 12)
	fmt.Println("Sum of 18, 8, 9, 10, 12 is", proResult)
	fmt.Println("Pro Message is", proMessage)
}

// func birthdayWish() {
// 	fmt.Println("Happy Birthday")
// }

// func anniversaryWish() {
// 	fmt.Println("Happy Anniversary")
// }

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func proSum(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Returning a message"
}
