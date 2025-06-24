package main

import "fmt"

func main() {
	fmt.Println("welcome to Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for range loop
	for index, day := range days {
		fmt.Printf("Index: %d, Day: %s\n", index, day)
	}

	// while loop equivalent in Go
	i := 0
	for i < len(days) {
		fmt.Println(days[i])
		i++
	}
}
