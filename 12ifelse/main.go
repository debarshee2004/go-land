package main

import "fmt"

func main() {
	fmt.Println("Welcome to the if else section of Go programming!")

	age := 20
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
}
