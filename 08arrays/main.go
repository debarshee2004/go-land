package main

import "fmt"

func main() {
	fmt.Println("Welcome to the arrays section of Go programming!")

	// 1D array
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"
	fruits[3] = "Date"
	fmt.Println("Fruits array:", fruits)
	fmt.Println("Length of fruits array:", len(fruits))

	// 2D array
	var vegetables [2][3]string
	vegetables[0][0] = "Carrot"
	vegetables[0][1] = "Broccoli"
	vegetables[0][2] = "Spinach"
	vegetables[1][0] = "Potato"
	vegetables[1][1] = "Onion"
	vegetables[1][2] = "Cabbage"
	fmt.Println("Vegetables array:", vegetables)
	fmt.Println("Length of vegetables array:", len(vegetables))
}
