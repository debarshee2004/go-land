package main

import "fmt"

func main() {
	fmt.Println("Welcome to the switch case section of Go programming!")

	day := 3

	switch day {
	case 1:
		fmt.Println("It's Monday!")
	case 2:
		fmt.Println("It's Tuesday!")
	case 3:
		fmt.Println("It's Wednesday!")
	case 4:
		fmt.Println("It's Thursday!")
	case 5:
		fmt.Println("It's Friday!")
	case 6, 7:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day!")
	}
	fmt.Println("Switch case example completed.")
}
