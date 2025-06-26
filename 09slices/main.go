package main

import "fmt"

func main() {
	fmt.Println("Welcome to the slices section of Go programming!")

	// ----------------------------
	// Example 1: Declaring and printing a 1D slice
	fruits := []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("\nExample 1: Initial Slice")
	fmt.Println("Fruits slice:", fruits)
	fmt.Printf("Length: %d, Capacity: %d\n", len(fruits), cap(fruits))
	fmt.Printf("Type of fruits: %T\n", fruits)

	// ----------------------------
	// Example 2: Appending elements
	fmt.Println("\nExample 2: Appending to a Slice")
	fruits = append(fruits, "Elderberry", "Grape")
	fmt.Println("Updated Fruits slice:", fruits)
	fmt.Printf("Length: %d, New Capacity: %d\n", len(fruits), cap(fruits))

	// ----------------------------
	// Example 3: Removing an element from a slice
	fmt.Println("\nExample 3: Removing an element from a slice")
	courses := []string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println("Original Courses slice:", courses)
	indexToRemove := 2 // Removes "Java"
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println("Courses after removal:", courses)

	// ----------------------------
	// Example 4: Creating a slice using make()
	fmt.Println("\nExample 4: Using make() to create slices")
	numbers := make([]int, 5, 10) // length = 5, capacity = 10
	fmt.Println("Slice created with make():", numbers)
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// ----------------------------
	// Example 5: Slicing a slice (sub-slicing)
	fmt.Println("\nExample 5: Sub-slicing")
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	sub := primes[2:5] // Elements at index 2 to 4
	fmt.Println("Original slice:", primes)
	fmt.Println("Sub-slice (2:5):", sub)

	// ----------------------------
	// Example 6: Appending a slice to another slice
	fmt.Println("\nExample 6: Appending one slice to another")
	even := []int{2, 4, 6}
	odd := []int{1, 3, 5}
	numbersCombined := append(even, odd...)
	fmt.Println("Combined slice:", numbersCombined)
}
