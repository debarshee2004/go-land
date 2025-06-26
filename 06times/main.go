package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Times package!")

	// ----------------------------
	// Example 1: Getting the current time
	presentTime := time.Now()
	fmt.Println("\nExample 1: Current Time")
	fmt.Println("Current time (raw):", presentTime)
	fmt.Printf("Type of presentTime: %T\n", presentTime)

	// ----------------------------
	// Example 2: Formatting time
	fmt.Println("\nExample 2: Formatting Time")
	// Layout must be "Mon Jan 2 15:04:05 MST 2006"
	formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday") // MM-DD-YYYY HH:MM:SS Day
	fmt.Println("Formatted time:", formattedTime)

	// ----------------------------
	// Example 3: Parsing time string into time.Time
	fmt.Println("\nExample 3: Parsing Time")
	parsedTime, err := time.Parse("01-02-2006 15:04:05 Monday", formattedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed time object:", parsedTime)
	}

	// ----------------------------
	// Example 4: Creating a specific time (e.g., for scheduling)
	fmt.Println("\nExample 4: Creating a Specific Time")
	customTime := time.Date(2023, time.December, 25, 14, 30, 0, 0, time.UTC)
	fmt.Println("Custom created time:", customTime)
	fmt.Printf("Type of customTime: %T\n", customTime)

	// ----------------------------
	// Example 5: Duration between two times
	fmt.Println("\nExample 5: Calculating Time Duration")
	start := time.Now()
	time.Sleep(2 * time.Second) // Simulate delay
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Time elapsed: %v\n", duration)

	// ----------------------------
	// Example 6: Adding/Subtracting time
	fmt.Println("\nExample 6: Time Arithmetic")
	tomorrow := presentTime.Add(24 * time.Hour)
	fmt.Println("Tomorrow will be:", tomorrow)

	yesterday := presentTime.Add(-24 * time.Hour)
	fmt.Println("Yesterday was:", yesterday)
}
