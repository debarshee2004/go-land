package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Functions section of Go programming!")

	// ----------------------------
	// Example 1: Basic function call
	fmt.Println("\nExample 1: Basic function call")
	greet()

	// ----------------------------
	// Example 2: Function with parameters
	fmt.Println("\nExample 2: Function with parameters")
	printSum(10, 20)

	// ----------------------------
	// Example 3: Function with return value
	fmt.Println("\nExample 3: Function with return value")
	result := multiply(4, 5)
	fmt.Println("Multiplication Result:", result)

	// ----------------------------
	// Example 4: Function with multiple return values
	fmt.Println("\nExample 4: Function with multiple return values")
	name, length := formatName("debarshee")
	fmt.Printf("Formatted Name: %s, Length: %d\n", name, length)

	// ----------------------------
	// Example 5: Named return values
	fmt.Println("\nExample 5: Named return values")
	diff := subtract(20, 7)
	fmt.Println("Difference:", diff)

	// ----------------------------
	// Example 6: Variadic function (takes variable number of args)
	fmt.Println("\nExample 6: Variadic function")
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", total)

	// ----------------------------
	// Example 7: Anonymous function (function literal)
	fmt.Println("\nExample 7: Anonymous function")
	func(msg string) {
		fmt.Println("Anonymous function says:", msg)
	}("Hi from inline!")

	// ----------------------------
	// Example 8: Assign function to a variable
	fmt.Println("\nExample 8: Function assigned to variable")
	double := func(n int) int {
		return n * 2
	}
	fmt.Println("Double of 6 is:", double(6))
}

// ----------------------------
// Example 1: Basic function
func greet() {
	fmt.Println("Hello from greet()!")
}

// Example 2: Function with parameters
func printSum(a int, b int) {
	fmt.Printf("Sum of %d and %d is %d\n", a, b, a+b)
}

// Example 3: Function with return value
func multiply(x int, y int) int {
	return x * y
}

// Example 4: Function with multiple return values
func formatName(name string) (string, int) {
	formatted := strings.Title(name) // Capitalizes first letter
	length := len(formatted)
	return formatted, length
}

// Example 5: Named return values
func subtract(a int, b int) (difference int) {
	difference = a - b
	return // returns `difference` implicitly
}

// Example 6: Variadic function
func sumAll(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}
