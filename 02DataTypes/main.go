// Data Type

package main

import "fmt"

const loginToken string = "Akhtar"

func main() {
	// String Data Type
	var username string = "Shadab"
	fmt.Println(username)
	fmt.Printf("Variable Type is: %T \n", username)

	// Bool Data Type
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable Type is: %T \n", isLoggedIn)

	// uint8 Data Type
	var smallValueInteger uint8 = 255
	fmt.Println(smallValueInteger)
	fmt.Printf("Variable Type is: %T \n", smallValueInteger)

	// uint16 Data Type
	var largeValueInteger uint16 = 299
	fmt.Println(largeValueInteger)
	fmt.Printf("Variable Type is: %T \n", largeValueInteger)

	// float32 Data Type
	var smallValuefloat float32 = 299.9467363527847
	fmt.Println(smallValuefloat)
	fmt.Printf("Variable Type is: %T \n", smallValuefloat)

	// float64 Data Type
	var bigValuefloat float64 = 299.9467363527847
	fmt.Println(bigValuefloat)
	fmt.Printf("Variable Type is: %T \n", bigValuefloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables Type is: %T \n", anotherVariable)

	var anotherVariableString string
	fmt.Println(anotherVariableString)
	fmt.Printf("Variables Type is: %T \n", anotherVariableString)

	// implicit Type
	var website = "skillwise.com"
	// fmt.Println(website)
	// fmt.Printf("Variables Type is: %T \n", website)

	// no var keyword declaration
	totalUsers := 49736
	fmt.Println(totalUsers)
	fmt.Printf("Variables Type is: %T \n", totalUsers)

	// const
	fmt.Println(loginToken)
	fmt.Printf("Variables Type is: %T \n", loginToken)

}
