package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the maps section of Go programming!")

	// ----------------------------
	// Example 1: Creating a map using make()
	languages := make(map[string]string) // key: string, value: string

	// Adding key-value pairs
	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"
	languages["de"] = "German"
	languages["zh"] = "Chinese"

	fmt.Println("\nExample 1: Initial map contents")
	fmt.Println("Languages map:", languages)

	// ----------------------------
	// Example 2: Accessing a value
	fmt.Println("\nExample 2: Accessing values")
	code := "fr"
	fmt.Printf("Language with code '%s': %s\n", code, languages[code])

	// ----------------------------
	// Example 3: Iterating through a map
	fmt.Println("\nExample 3: Looping through map entries")
	for code, name := range languages {
		fmt.Printf("Language code: %s, Language name: %s\n", code, name)
	}

	// ----------------------------
	// Example 4: Deleting an entry
	fmt.Println("\nExample 4: Deleting a key")
	delete(languages, "zh") // Removes Chinese
	fmt.Println("Map after deleting key 'zh':", languages)

	// ----------------------------
	// Example 5: Check if a key exists
	fmt.Println("\nExample 5: Checking if a key exists")
	val, exists := languages["en"]
	if exists {
		fmt.Printf("'en' exists with value: %s\n", val)
	} else {
		fmt.Println("'en' does not exist.")
	}

	_, exists2 := languages["jp"]
	if !exists2 {
		fmt.Println("'jp' does not exist in the map.")
	}

	// ----------------------------
	// Example 6: Declaring and initializing a map in one line
	fmt.Println("\nExample 6: Map literal (inline initialization)")
	currencies := map[string]string{
		"INR": "Indian Rupee",
		"USD": "US Dollar",
		"EUR": "Euro",
	}
	fmt.Println("Currency map:", currencies)
}
