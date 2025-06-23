package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the maps section of Go programming!")
	languages := make(map[string]string)
	languages["en"] = "English"
	languages["es"] = "Spanish"
	languages["fr"] = "French"
	languages["de"] = "German"
	languages["zh"] = "Chinese"

	fmt.Println("Languages map:", languages)

	// Looping through the map
	// keys: code, values: name
	for code, name := range languages {
		fmt.Printf("Language code: %s, Language name: %s\n", code, name)
	}

	// Deleting an entry
	delete(languages, "zh")
	fmt.Println("After deletion, languages map:", languages)
}
