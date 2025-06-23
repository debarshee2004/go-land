package main

import "fmt"

func main() {
	fmt.Println("Welcome to the slices section of Go programming!")

	// 1D slice
	// fruits := []string{"Apple", "Banana", "Cherry", "Date"}
	// fmt.Println("Fruits slice:", fruits)
	// fmt.Printf("Capacity of fruits slice: %d\n", cap(fruits))
	// fmt.Printf("Type of fruits slice: %T\n", fruits)

	// fruits = append(fruits, "Elderberry", "Grape")
	// fmt.Println("Updated fruits slice:", fruits)
	// fmt.Printf("New capacity of fruits slice: %d\n", cap(fruits))

	// how ro remove a value from slices based on index
	cources := []string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println("Courses slice before removal:", cources)
	indexToRemove := 2 // Remove "Java"
	cources = append(cources[:indexToRemove], cources[indexToRemove+1:]...)
	fmt.Println("Courses slice after removal:", cources)
}
