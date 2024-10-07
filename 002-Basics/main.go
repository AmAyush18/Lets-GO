package main

import "fmt"

func main() {
	variables() // variables in GO
}

const LoginToken string = "a12sda2df1a3fdas5f46adas64" // having first letter of variable_name as capital we define it as public
// numberOfUser := 300000  -> varless operator not allowed outside methods

func variables() {
	// string
	var username string = "amayush18"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// integers
	var smallVal uint8 = 255 // not pos
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// floats
	var smallFloat float32 = 255.546464465423146465413 // not pos
	var longFloat float64 = 255.546464465423146465413  // not pos
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	fmt.Println(longFloat)
	fmt.Printf("Variable is of type: %T \n", longFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "https://ayush-dwivedi.netlify.app" // -> it will be given a type based on value we pass it
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	var number = 3
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)
	// website = 3  -> Not allowed since lexer has now considered it to be string

	// no var style
	numberOfUser := 300000 // varless operator (using := and not using var) is allowed inside method but not outside method
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
