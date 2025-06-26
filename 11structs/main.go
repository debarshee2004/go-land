package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println("Welcome to the structs section of Go programming!")

	// ----------------------------
	// Example 1: Defining a struct

	// Creating an instance using field names
	alice := Person{
		FirstName: "Alice",
		LastName:  "Johnson",
		Age:       30,
	}

	// Accessing fields
	fmt.Println("\nExample 1: Struct field access")
	fmt.Printf("Name: %s %s, Age: %d\n", alice.FirstName, alice.LastName, alice.Age)

	// Updating a field
	alice.Age = 31
	fmt.Printf("Updated Age: %d\n", alice.Age)

	// ----------------------------
	// Example 2: Using anonymous struct
	fmt.Println("\nExample 2: Anonymous Struct")
	book := struct {
		Title  string
		Author string
		Pages  int
	}{
		Title:  "Go in Action",
		Author: "William Kennedy",
		Pages:  300,
	}
	fmt.Printf("Book: %+v\n", book)

	// ----------------------------
	// Example 3: Pointer to a struct
	fmt.Println("\nExample 3: Pointer to Struct")
	bob := &Person{
		FirstName: "Bob",
		LastName:  "Smith",
		Age:       28,
	}
	bob.Age++ // Modify via pointer
	fmt.Printf("Bob's details (via pointer): %+v\n", bob)

	// ----------------------------
	// Example 4: Slice of structs
	fmt.Println("\nExample 4: Slice of Structs")
	people := []Person{
		{FirstName: "Charlie", LastName: "Brown", Age: 22},
		{FirstName: "Daisy", LastName: "Hill", Age: 27},
	}
	for _, p := range people {
		fmt.Printf("Person: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
	}

	// ----------------------------
	// Example 5: Struct with method
	fmt.Println("\nExample 5: Struct with Method")

	// Method defined below main()
	john := Person{"John", "Doe", 35}
	john.Greet()
}

// ----------------------------
// Method attached to Person struct
func (p Person) Greet() {
	fmt.Printf("Hello, I'm %s %s and I'm %d years old.\n", p.FirstName, p.LastName, p.Age)
}
