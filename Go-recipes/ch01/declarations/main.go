package main

import (
	"fmt"
)

// Declare constant
const Title = "Person Details"

// Declare package variable
var Country = "USA"

func main() {
	// Print constant variable
	fmt.Println(Title)

	fname, lname := "Shiju", "Varghese"
	age := 35

	// Print local variables
	fmt.Println("First Name:", fname)
	fmt.Println("Last Name:", lname)
	fmt.Println("Age:", age)

	fmt.Println("First Name:", fname, ",Last Name:", lname, ",Age:", age)
	// Print package variable

	Country = "China"
	fmt.Println("Country:", Country)

	Country := "China"
	fmt.Println("Country:", Country)

	//Title = "My Details" //常量不可以修改
	Title := "My Details"
	fmt.Println(Title)
}
