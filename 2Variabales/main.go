package main

import "fmt"

var PublicVar = 123

func main() {
	var username string = "Anurag"
	var smallVal uint8 = 255
	var floatr float64 = 255.3434342323232323232323
	fmt.Println(smallVal)
	fmt.Println(username)
	fmt.Println(floatr)
	fmt.Printf("Variable is of type:  %T \n", username)

	// default value
	var anotherVar int
	fmt.Println(anotherVar)

	// implicit type
	var website = "villagetilpat.in"
	fmt.Println(website)
	// no var style onlt in methods
	name1 := "Anurag"
	fmt.Println(name1)

	fmt.Println(PublicVar)
}
