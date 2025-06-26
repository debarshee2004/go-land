package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Example 1: Basic string input from user
	fmt.Println("Example 1: String Input")
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)  // Create a buffered reader to read input from stdin
	name, err := reader.ReadString('\n') // Read input until newline (\n)
	if err != nil {                      // Error handling
		fmt.Println("Error reading input:", err)
	} else {
		name = strings.TrimSpace(name) // Trim newline characters
		fmt.Println("Hello,", name)
	}

	// Example 2: Integer input
	fmt.Println("\nExample 2: Integer Input")
	fmt.Print("Enter your age: ")
	ageInput, _ := reader.ReadString('\n') // Reuse reader to read age input
	ageInput = strings.TrimSpace(ageInput) // Clean input string
	age, err := strconv.Atoi(ageInput)     // Convert string to int
	if err != nil {
		fmt.Println("Invalid age input:", err)
	} else {
		fmt.Println("You are", age, "years old.")
	}

	// Example 3: Multiple values in one line
	fmt.Println("\nExample 4: Reading Multiple Values in One Line")
	fmt.Print("Enter your city and country (e.g., Kolkata India): ")
	locationInput, _ := reader.ReadString('\n')
	locationInput = strings.TrimSpace(locationInput)
	parts := strings.Split(locationInput, " ") // Split by space
	if len(parts) >= 2 {
		city := parts[0]
		country := parts[1]
		fmt.Printf("City: %s, Country: %s\n", city, country)
	} else {
		fmt.Println("Please provide both city and country.")
	}

	// Example 4: Using fmt.Scan for simple inputs (alternative method)
	fmt.Println("\nExample 5: Using fmt.Scan")
	var favLang string
	fmt.Print("Enter your favorite programming language: ")
	fmt.Scan(&favLang) // Directly scans input until space or newline
	fmt.Println("You love", favLang)
}
