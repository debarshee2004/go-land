package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs section of Go programming!")

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	// Creating an instance of Person
	alice := Person{
		FirstName: "Alice",
		LastName:  "Johnson",
		Age:       30,
	}

	// Accessing fields
	fmt.Printf("Name: %s %s, Age: %d\n", alice.FirstName, alice.LastName, alice.Age)

	// Updating a field
	alice.Age = 31
	fmt.Printf("Updated Age: %d\n", alice.Age)
}
