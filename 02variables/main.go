package main

import (
	"fmt"
)

func main() {
	// Declare a variable of type int
	var age uint = 30
	fmt.Println("Age:", age)
	fmt.Printf("Age is of type: %T \n", age)

	// Declare a variable of type string
	var name string = "Alice"
	fmt.Println("Name:", name)
	fmt.Printf("Name is of type: %T \n", name)

	// Declare a variable of type float64
	var height float64 = 5.7
	fmt.Println("Height:", height)
	fmt.Printf("Height is of type: %T \n", height)

	// Declare a variable of type bool
	var isStudent bool = false
	fmt.Println("Is Student:", isStudent)
	fmt.Printf("Is Student is of type: %T \n", isStudent)

	// ------------------------------------

	// Declare a variable using shorthand notation
	country := "USA"
	fmt.Println("Country:", country)
	fmt.Printf("Country is of type: %T \n", country)

	number := 123
	fmt.Println("Number:", number)
	fmt.Printf("Number is of type: %T \n", number)

	pi := 3.14
	fmt.Println("Pi:", pi)
	fmt.Printf("Pi is of type: %T \n", pi)

	isHungry := true
	fmt.Println("Is Hungry:", isHungry)
	fmt.Printf("Is Hungry is of type: %T \n", isHungry)
}
