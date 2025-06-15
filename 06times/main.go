package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Times package!")

	presentTime := time.Now()
	fmt.Println("Current time:", presentTime)
	fmt.Printf("Current time type: %T\n", presentTime)

	// Formatting time
	formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Formatted time:", formattedTime)

	// Parsing time
	parsedTime, err := time.Parse("01-02-2006 15:04:05 Monday", formattedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed time:", parsedTime)
	}

	// Creating a specific time
	createTime := time.Date(2023, time.December, 25, 14, 30, 0, 0, time.UTC)
	fmt.Println("Created time:", createTime)
	fmt.Printf("Created time type: %T\n", createTime)
}
